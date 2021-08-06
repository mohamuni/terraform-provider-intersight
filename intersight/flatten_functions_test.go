package intersight

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func compareMaps(expected, ffOutput map[string]interface{}, t *testing.T) error {
	for key, value := range expected {
		v, ok := ffOutput[key]
		if !ok {
			return fmt.Errorf("error: key not found %s", key)
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			value := value.([]map[string]interface{})
			v := v.([]map[string]interface{})
			for i := 0; i < len(value); i++ {
				if err := compareMaps(value[i], v[i], t); err != nil {
					return err
				}
			}
		case reflect.Map:
			return compareMaps(value.(map[string]interface{}), v.(map[string]interface{}), t)
		case reflect.Int32, reflect.Int, reflect.Int64:
			if reflect.ValueOf(value).Int() != reflect.ValueOf(v).Int() {
				return fmt.Errorf("error: mismatch in int type: expectedValue %+v kind %T, gotValue %+v kind %T", value, value, v,
					v)
			}
		case reflect.Float64, reflect.Float32:
			if reflect.ValueOf(value).Float() != reflect.ValueOf(v).Float() {
				return fmt.Errorf("error: mismatch in float type: expectedValue %+v kind %T, gotValue %+v kind %T", value, value, v,
					v)
			}
		default:
			if !reflect.DeepEqual(value, v) {
				return fmt.Errorf("error: mismatch in value %+v %+v", v, value)
			}
		}
	}
	return nil
}

//CheckError Panics when error occurs
func CheckError(t *testing.T, e error) {
	if e != nil {
		t.Errorf("Error while testing flatten function: %+v", e)
	}
}

func TestFlattenListAdapterAdapterConfig(t *testing.T) {
	p := []models.AdapterAdapterConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"adapter.AdapterConfig","ObjectType":"adapter.AdapterConfig","SlotId":"SlotId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterAdapterConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterAdapterConfig{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterAdapterConfig(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "adapter.AdapterConfig", "object_type": "adapter.AdapterConfig", "slot_id": "SlotId 1"}, {"class_id": "adapter.AdapterConfig", "object_type": "adapter.AdapterConfig", "slot_id": "SlotId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterExtEthInterfaceRelationship(t *testing.T) {
	p := []models.AdapterExtEthInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterExtEthInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterExtEthInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterExtEthInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostEthInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostEthInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostEthInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostEthInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostEthInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostFcInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostFcInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostFcInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostFcInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostFcInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostIscsiInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostIscsiInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostIscsiInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostIscsiInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostIscsiInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterUnitRelationship(t *testing.T) {
	p := []models.AdapterUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceApiStatus(t *testing.T) {
	p := []models.ApplianceApiStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"appliance.ApiStatus","ObjectType":"appliance.ApiStatus","ElapsedTime":32.000000,"ObjectName":"ObjectName %d","Reason":"Reason %d","Response":32}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceApiStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceApiStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceApiStatus(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "appliance.ApiStatus", "object_type": "appliance.ApiStatus", "elapsed_time": 32.000000, "object_name": "ObjectName 1", "reason": "Reason 1", "response": 32}, {"class_id": "appliance.ApiStatus", "object_type": "appliance.ApiStatus", "elapsed_time": 32.000000, "object_name": "ObjectName 2", "reason": "Reason 2", "response": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceAppStatusRelationship(t *testing.T) {
	p := []models.ApplianceAppStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceAppStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceAppStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceAppStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceCertRenewalPhase(t *testing.T) {
	p := []models.ApplianceCertRenewalPhase{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"appliance.CertRenewalPhase","ClassId":"appliance.CertRenewalPhase","Name":"Name %d","Failed":true,"Message":"Message %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceCertRenewalPhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceCertRenewalPhase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceCertRenewalPhase(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "appliance.CertRenewalPhase", "class_id": "appliance.CertRenewalPhase", "name": "Name 1", "failed": true, "message": "Message 1"}, {"object_type": "appliance.CertRenewalPhase", "class_id": "appliance.CertRenewalPhase", "name": "Name 2", "failed": true, "message": "Message 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceDataExportPolicyRelationship(t *testing.T) {
	p := []models.ApplianceDataExportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceDataExportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceDataExportPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceDataExportPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceFileSystemStatusRelationship(t *testing.T) {
	p := []models.ApplianceFileSystemStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceFileSystemStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceFileSystemStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceFileSystemStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceGroupStatusRelationship(t *testing.T) {
	p := []models.ApplianceGroupStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceGroupStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceGroupStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceGroupStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceKeyValuePair(t *testing.T) {
	p := []models.ApplianceKeyValuePair{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"appliance.KeyValuePair","ObjectType":"appliance.KeyValuePair","Value":"Value %d","Key":"Key %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceKeyValuePair(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceKeyValuePair{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceKeyValuePair(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "appliance.KeyValuePair", "object_type": "appliance.KeyValuePair", "value": "Value 1", "key": "Key 1"}, {"class_id": "appliance.KeyValuePair", "object_type": "appliance.KeyValuePair", "value": "Value 2", "key": "Key 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceStatusCheck(t *testing.T) {
	p := []models.ApplianceStatusCheck{}
	var d = &schema.ResourceData{}
	c := `{"Code":"Code %d","Result":"Result %d","ClassId":"appliance.StatusCheck","ObjectType":"appliance.StatusCheck"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceStatusCheck(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceStatusCheck{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceStatusCheck(p, d)
	expectedOp := []map[string]interface{}{{"code": "Code 1", "result": "Result 1", "class_id": "appliance.StatusCheck", "object_type": "appliance.StatusCheck"}, {"code": "Code 2", "result": "Result 2", "class_id": "appliance.StatusCheck", "object_type": "appliance.StatusCheck"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetClusterMemberRelationship(t *testing.T) {
	p := []models.AssetClusterMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetClusterMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetClusterMemberRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetClusterMemberRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetConnection(t *testing.T) {
	p := []models.AssetConnection{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"asset.Connection","ClassId":"asset.Connection"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetConnection(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetConnection{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetConnection(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "asset.Connection", "class_id": "asset.Connection"}, {"object_type": "asset.Connection", "class_id": "asset.Connection"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeploymentRelationship(t *testing.T) {
	p := []models.AssetDeploymentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeploymentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeploymentRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeploymentRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeploymentDeviceRelationship(t *testing.T) {
	p := []models.AssetDeploymentDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeploymentDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeploymentDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeploymentDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeviceRegistrationRelationship(t *testing.T) {
	p := []models.AssetDeviceRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeviceRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeviceRegistrationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeviceRegistrationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetMeteringType(t *testing.T) {
	p := []models.AssetMeteringType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.MeteringType","ObjectType":"asset.MeteringType","Name":"Name %d","Unit":"Unit %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetMeteringType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetMeteringType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetMeteringType(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "asset.MeteringType", "object_type": "asset.MeteringType", "name": "Name 1", "unit": "Unit 1"}, {"class_id": "asset.MeteringType", "object_type": "asset.MeteringType", "name": "Name 2", "unit": "Unit 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetService(t *testing.T) {
	p := []models.AssetService{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"asset.Service","ClassId":"asset.Service"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetService(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetService{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetService(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "asset.Service", "class_id": "asset.Service"}, {"object_type": "asset.Service", "class_id": "asset.Service"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBiosBootDeviceRelationship(t *testing.T) {
	p := []models.BiosBootDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBiosBootDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BiosBootDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBiosBootDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBiosUnitRelationship(t *testing.T) {
	p := []models.BiosUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBiosUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BiosUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBiosUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootCddDeviceRelationship(t *testing.T) {
	p := []models.BootCddDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootCddDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootCddDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootCddDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootDeviceBase(t *testing.T) {
	p := []models.BootDeviceBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"boot.DeviceBase","Enabled":true,"Name":"Name %d","ObjectType":"boot.DeviceBase"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootDeviceBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootDeviceBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootDeviceBase(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "boot.DeviceBase", "enabled": true, "name": "Name 1", "object_type": "boot.DeviceBase"}, {"class_id": "boot.DeviceBase", "enabled": true, "name": "Name 2", "object_type": "boot.DeviceBase"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootHddDeviceRelationship(t *testing.T) {
	p := []models.BootHddDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootHddDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootHddDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootHddDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootIscsiDeviceRelationship(t *testing.T) {
	p := []models.BootIscsiDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootIscsiDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootIscsiDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootIscsiDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootNvmeDeviceRelationship(t *testing.T) {
	p := []models.BootNvmeDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootNvmeDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootNvmeDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootNvmeDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootPchStorageDeviceRelationship(t *testing.T) {
	p := []models.BootPchStorageDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootPchStorageDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootPchStorageDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootPchStorageDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootPxeDeviceRelationship(t *testing.T) {
	p := []models.BootPxeDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootPxeDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootPxeDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootPxeDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootSanDeviceRelationship(t *testing.T) {
	p := []models.BootSanDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootSanDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootSanDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootSanDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootSdDeviceRelationship(t *testing.T) {
	p := []models.BootSdDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootSdDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootSdDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootSdDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootUefiShellDeviceRelationship(t *testing.T) {
	p := []models.BootUefiShellDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootUefiShellDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootUefiShellDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootUefiShellDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootUsbDeviceRelationship(t *testing.T) {
	p := []models.BootUsbDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootUsbDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootUsbDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootUsbDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootVmediaDeviceRelationship(t *testing.T) {
	p := []models.BootVmediaDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootVmediaDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootVmediaDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootVmediaDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilityCapabilityRelationship(t *testing.T) {
	p := []models.CapabilityCapabilityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilityCapabilityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilityCapabilityRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilityCapabilityRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilityPortRange(t *testing.T) {
	p := []models.CapabilityPortRange{}
	var d = &schema.ResourceData{}
	c := `{"StartSlotId":32,"EndPortId":32,"EndSlotId":32,"ClassId":"capability.PortRange","ObjectType":"capability.PortRange","StartPortId":32}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilityPortRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilityPortRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilityPortRange(p, d)
	expectedOp := []map[string]interface{}{{"start_slot_id": 32, "end_port_id": 32, "end_slot_id": 32, "class_id": "capability.PortRange", "object_type": "capability.PortRange", "start_port_id": 32}, {"start_slot_id": 32, "end_port_id": 32, "end_slot_id": 32, "class_id": "capability.PortRange", "object_type": "capability.PortRange", "start_port_id": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilitySwitchingModeCapability(t *testing.T) {
	p := []models.CapabilitySwitchingModeCapability{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"capability.SwitchingModeCapability","SwitchingMode":"SwitchingMode %d","VpCompressionSupported":true,"ClassId":"capability.SwitchingModeCapability"}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilitySwitchingModeCapability(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilitySwitchingModeCapability{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilitySwitchingModeCapability(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "capability.SwitchingModeCapability", "switching_mode": "SwitchingMode 1", "vp_compression_supported": true, "class_id": "capability.SwitchingModeCapability"}, {"object_type": "capability.SwitchingModeCapability", "switching_mode": "SwitchingMode 2", "vp_compression_supported": true, "class_id": "capability.SwitchingModeCapability"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCertificatemanagementCertificateBase(t *testing.T) {
	p := []models.CertificatemanagementCertificateBase{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"IsPrivatekeySet":true,"ObjectType":"certificatemanagement.CertificateBase","ClassId":"certificatemanagement.CertificateBase"}`

	//test when the response is empty
	ffOpEmpty := flattenListCertificatemanagementCertificateBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CertificatemanagementCertificateBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCertificatemanagementCertificateBase(p, d)
	expectedOp := []map[string]interface{}{{"enabled": true, "is_privatekey_set": true, "object_type": "certificatemanagement.CertificateBase", "class_id": "certificatemanagement.CertificateBase"}, {"enabled": true, "is_privatekey_set": true, "object_type": "certificatemanagement.CertificateBase", "class_id": "certificatemanagement.CertificateBase"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisConfigChangeDetailRelationship(t *testing.T) {
	p := []models.ChassisConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisConfigResultEntryRelationship(t *testing.T) {
	p := []models.ChassisConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisIomProfileRelationship(t *testing.T) {
	p := []models.ChassisIomProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisIomProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisIomProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisIomProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudCustomAttributes(t *testing.T) {
	p := []models.CloudCustomAttributes{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.CustomAttributes","AttributeName":"AttributeName %d","AttributeType":"AttributeType %d","AttributeValue":"AttributeValue %d","ObjectType":"cloud.CustomAttributes"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudCustomAttributes(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudCustomAttributes{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudCustomAttributes(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "cloud.CustomAttributes", "attribute_name": "AttributeName 1", "attribute_type": "AttributeType 1", "attribute_value": "AttributeValue 1", "object_type": "cloud.CustomAttributes"}, {"class_id": "cloud.CustomAttributes", "attribute_name": "AttributeName 2", "attribute_type": "AttributeType 2", "attribute_value": "AttributeValue 2", "object_type": "cloud.CustomAttributes"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudTfcWorkspaceVariables(t *testing.T) {
	p := []models.CloudTfcWorkspaceVariables{}
	var d = &schema.ResourceData{}
	c := `{"Hcl":true,"ClassId":"cloud.TfcWorkspaceVariables","Identity":"Identity %d","Sensitive":true,"Category":"Category %d","Key":"Key %d","ObjectType":"cloud.TfcWorkspaceVariables","Description":"Description %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudTfcWorkspaceVariables(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudTfcWorkspaceVariables{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudTfcWorkspaceVariables(p, d)
	expectedOp := []map[string]interface{}{{"hcl": true, "class_id": "cloud.TfcWorkspaceVariables", "identity": "Identity 1", "sensitive": true, "category": "Category 1", "key": "Key 1", "object_type": "cloud.TfcWorkspaceVariables", "description": "Description 1", "value": "Value 1"}, {"hcl": true, "class_id": "cloud.TfcWorkspaceVariables", "identity": "Identity 2", "sensitive": true, "category": "Category 2", "key": "Key 2", "object_type": "cloud.TfcWorkspaceVariables", "description": "Description 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeBladeRelationship(t *testing.T) {
	p := []models.ComputeBladeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeBladeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeBladeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeBladeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeIpAddress(t *testing.T) {
	p := []models.ComputeIpAddress{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"compute.IpAddress","Dn":"Dn %d","KvmVlan":32,"Name":"Name %d","HttpPort":32,"HttpsPort":32,"Category":"Category %d","DefaultGateway":"DefaultGateway %d","KvmPort":32,"ClassId":"compute.IpAddress","Subnet":"Subnet %d","Type":"Type %d","Address":"Address %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeIpAddress(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeIpAddress{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeIpAddress(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "compute.IpAddress", "dn": "Dn 1", "kvm_vlan": 32, "name": "Name 1", "http_port": 32, "https_port": 32, "category": "Category 1", "default_gateway": "DefaultGateway 1", "kvm_port": 32, "class_id": "compute.IpAddress", "subnet": "Subnet 1", "type": "Type 1", "address": "Address 1"}, {"object_type": "compute.IpAddress", "dn": "Dn 2", "kvm_vlan": 32, "name": "Name 2", "http_port": 32, "https_port": 32, "category": "Category 2", "default_gateway": "DefaultGateway 2", "kvm_port": 32, "class_id": "compute.IpAddress", "subnet": "Subnet 2", "type": "Type 2", "address": "Address 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeMappingRelationship(t *testing.T) {
	p := []models.ComputeMappingRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeMappingRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeMappingRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeMappingRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputePhysicalRelationship(t *testing.T) {
	p := []models.ComputePhysicalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputePhysicalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputePhysicalRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputePhysicalRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeRackUnitRelationship(t *testing.T) {
	p := []models.ComputeRackUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeRackUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeRackUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeRackUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCondHclStatusDetailRelationship(t *testing.T) {
	p := []models.CondHclStatusDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListCondHclStatusDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CondHclStatusDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCondHclStatusDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListConfigExportedItemRelationship(t *testing.T) {
	p := []models.ConfigExportedItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListConfigExportedItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ConfigExportedItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListConfigExportedItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListConfigImportedItemRelationship(t *testing.T) {
	p := []models.ConfigImportedItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListConfigImportedItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ConfigImportedItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListConfigImportedItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListConfigMoRef(t *testing.T) {
	p := []models.ConfigMoRef{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","ClassId":"config.MoRef","ObjectType":"config.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListConfigMoRef(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ConfigMoRef{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListConfigMoRef(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "class_id": "config.MoRef", "object_type": "config.MoRef"}, {"moid": "Moid 2", "class_id": "config.MoRef", "object_type": "config.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListConnectorpackConnectorPackUpdate(t *testing.T) {
	p := []models.ConnectorpackConnectorPackUpdate{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","NewVersion":"NewVersion %d","ClassId":"connectorpack.ConnectorPackUpdate","ObjectType":"connectorpack.ConnectorPackUpdate","CurrentVersion":"CurrentVersion %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListConnectorpackConnectorPackUpdate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ConnectorpackConnectorPackUpdate{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListConnectorpackConnectorPackUpdate(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "new_version": "NewVersion 1", "class_id": "connectorpack.ConnectorPackUpdate", "object_type": "connectorpack.ConnectorPackUpdate", "current_version": "CurrentVersion 1"}, {"name": "Name 2", "new_version": "NewVersion 2", "class_id": "connectorpack.ConnectorPackUpdate", "object_type": "connectorpack.ConnectorPackUpdate", "current_version": "CurrentVersion 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListContentComplexType(t *testing.T) {
	p := []models.ContentComplexType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"content.ComplexType","ObjectType":"content.ComplexType","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListContentComplexType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ContentComplexType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListContentComplexType(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "content.ComplexType", "object_type": "content.ComplexType", "name": "Name 1"}, {"class_id": "content.ComplexType", "object_type": "content.ComplexType", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListContentParameter(t *testing.T) {
	p := []models.ContentParameter{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"content.Parameter","AcceptSingleValue":true,"ObjectType":"content.Parameter","Secure":true,"ItemType":"ItemType %d","ComplexType":"ComplexType %d","Type":"Type %d","Path":"Path %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListContentParameter(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ContentParameter{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListContentParameter(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "content.Parameter", "accept_single_value": true, "object_type": "content.Parameter", "secure": true, "item_type": "ItemType 1", "complex_type": "ComplexType 1", "type": "Type 1", "path": "Path 1", "name": "Name 1"}, {"class_id": "content.Parameter", "accept_single_value": true, "object_type": "content.Parameter", "secure": true, "item_type": "ItemType 2", "complex_type": "ComplexType 2", "type": "Type 2", "path": "Path 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCrdCustomResourceConfigProperty(t *testing.T) {
	p := []models.CrdCustomResourceConfigProperty{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"crd.CustomResourceConfigProperty","ObjectType":"crd.CustomResourceConfigProperty","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCrdCustomResourceConfigProperty(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CrdCustomResourceConfigProperty{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCrdCustomResourceConfigProperty(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "crd.CustomResourceConfigProperty", "object_type": "crd.CustomResourceConfigProperty", "key": "Key 1", "value": "Value 1"}, {"class_id": "crd.CustomResourceConfigProperty", "object_type": "crd.CustomResourceConfigProperty", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentFanRelationship(t *testing.T) {
	p := []models.EquipmentFanRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentFanRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentFanRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentFanRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentFanModuleRelationship(t *testing.T) {
	p := []models.EquipmentFanModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentFanModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentFanModuleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentFanModuleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoCardRelationship(t *testing.T) {
	p := []models.EquipmentIoCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoCardIdentity(t *testing.T) {
	p := []models.EquipmentIoCardIdentity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"equipment.IoCardIdentity","ObjectType":"equipment.IoCardIdentity","IoCardMoid":"IoCardMoid %d","ModuleId":32,"NetworkElementMoid":"NetworkElementMoid %d","SwitchId":32}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoCardIdentity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoCardIdentity{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoCardIdentity(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "equipment.IoCardIdentity", "object_type": "equipment.IoCardIdentity", "io_card_moid": "IoCardMoid 1", "module_id": 32, "network_element_moid": "NetworkElementMoid 1", "switch_id": 32}, {"class_id": "equipment.IoCardIdentity", "object_type": "equipment.IoCardIdentity", "io_card_moid": "IoCardMoid 2", "module_id": 32, "network_element_moid": "NetworkElementMoid 2", "switch_id": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoExpanderRelationship(t *testing.T) {
	p := []models.EquipmentIoExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoExpanderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoExpanderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentPsuRelationship(t *testing.T) {
	p := []models.EquipmentPsuRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentPsuRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentPsuRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentPsuRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentRackEnclosureSlotRelationship(t *testing.T) {
	p := []models.EquipmentRackEnclosureSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentRackEnclosureSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentRackEnclosureSlotRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentRackEnclosureSlotRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentSwitchCardRelationship(t *testing.T) {
	p := []models.EquipmentSwitchCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentSwitchCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentSwitchCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentSwitchCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentSystemIoControllerRelationship(t *testing.T) {
	p := []models.EquipmentSystemIoControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentSystemIoControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentSystemIoControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentSystemIoControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentTpmRelationship(t *testing.T) {
	p := []models.EquipmentTpmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentTpmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentTpmRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentTpmRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherHostPortRelationship(t *testing.T) {
	p := []models.EtherHostPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherHostPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherHostPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherHostPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherNetworkPortRelationship(t *testing.T) {
	p := []models.EtherNetworkPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherNetworkPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherNetworkPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherNetworkPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherPhysicalPortRelationship(t *testing.T) {
	p := []models.EtherPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherPhysicalPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherPhysicalPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherPortChannelRelationship(t *testing.T) {
	p := []models.EtherPortChannelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherPortChannelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherPortChannelRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherPortChannelRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricConfigChangeDetailRelationship(t *testing.T) {
	p := []models.FabricConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricConfigResultEntryRelationship(t *testing.T) {
	p := []models.FabricConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricEthNetworkGroupPolicyRelationship(t *testing.T) {
	p := []models.FabricEthNetworkGroupPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricEthNetworkGroupPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricEthNetworkGroupPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricEthNetworkGroupPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricPortIdentifier(t *testing.T) {
	p := []models.FabricPortIdentifier{}
	var d = &schema.ResourceData{}
	c := `{"AggregatePortId":32,"PortId":32,"SlotId":32,"ClassId":"fabric.PortIdentifier","ObjectType":"fabric.PortIdentifier"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricPortIdentifier(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricPortIdentifier{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricPortIdentifier(p, d)
	expectedOp := []map[string]interface{}{{"aggregate_port_id": 32, "port_id": 32, "slot_id": 32, "class_id": "fabric.PortIdentifier", "object_type": "fabric.PortIdentifier"}, {"aggregate_port_id": 32, "port_id": 32, "slot_id": 32, "class_id": "fabric.PortIdentifier", "object_type": "fabric.PortIdentifier"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricQosClass(t *testing.T) {
	p := []models.FabricQosClass{}
	var d = &schema.ResourceData{}
	c := `{"Weight":32,"ClassId":"fabric.QosClass","Mtu":32,"BandwidthPercent":32,"Name":"Name %d","ObjectType":"fabric.QosClass","Cos":32,"AdminState":"AdminState %d","PacketDrop":true,"MulticastOptimize":true}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricQosClass(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricQosClass{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricQosClass(p, d)
	expectedOp := []map[string]interface{}{{"weight": 32, "class_id": "fabric.QosClass", "mtu": 32, "bandwidth_percent": 32, "name": "Name 1", "object_type": "fabric.QosClass", "cos": 32, "admin_state": "AdminState 1", "packet_drop": true, "multicast_optimize": true}, {"weight": 32, "class_id": "fabric.QosClass", "mtu": 32, "bandwidth_percent": 32, "name": "Name 2", "object_type": "fabric.QosClass", "cos": 32, "admin_state": "AdminState 2", "packet_drop": true, "multicast_optimize": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricSwitchProfileRelationship(t *testing.T) {
	p := []models.FabricSwitchProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricSwitchProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricSwitchProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricSwitchProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcPhysicalPortRelationship(t *testing.T) {
	p := []models.FcPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcPhysicalPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcPhysicalPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcPortChannelRelationship(t *testing.T) {
	p := []models.FcPortChannelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcPortChannelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcPortChannelRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcPortChannelRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcpoolBlock(t *testing.T) {
	p := []models.FcpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fcpool.Block","ObjectType":"fcpool.Block","Size":32,"From":"From %d","To":"To %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcpoolBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcpoolBlock(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "fcpool.Block", "object_type": "fcpool.Block", "size": 32, "from": "From 1", "to": "To 1"}, {"class_id": "fcpool.Block", "object_type": "fcpool.Block", "size": 32, "from": "From 2", "to": "To 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcpoolFcBlockRelationship(t *testing.T) {
	p := []models.FcpoolFcBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcpoolFcBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcpoolFcBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcpoolFcBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareBaseImpact(t *testing.T) {
	p := []models.FirmwareBaseImpact{}
	var d = &schema.ResourceData{}
	c := `{"DomainName":"DomainName %d","EndPoint":"EndPoint %d","VersionImpact":"VersionImpact %d","ClassId":"firmware.BaseImpact","ComputationError":"ComputationError %d","ComputationStatusDetail":"ComputationStatusDetail %d","FirmwareVersion":"FirmwareVersion %d","TargetFirmwareVersion":"TargetFirmwareVersion %d","ImpactType":"ImpactType %d","ObjectType":"firmware.BaseImpact","Model":"Model %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareBaseImpact(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareBaseImpact{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareBaseImpact(p, d)
	expectedOp := []map[string]interface{}{{"domain_name": "DomainName 1", "end_point": "EndPoint 1", "version_impact": "VersionImpact 1", "class_id": "firmware.BaseImpact", "computation_error": "ComputationError 1", "computation_status_detail": "ComputationStatusDetail 1", "firmware_version": "FirmwareVersion 1", "target_firmware_version": "TargetFirmwareVersion 1", "impact_type": "ImpactType 1", "object_type": "firmware.BaseImpact", "model": "Model 1"}, {"domain_name": "DomainName 2", "end_point": "EndPoint 2", "version_impact": "VersionImpact 2", "class_id": "firmware.BaseImpact", "computation_error": "ComputationError 2", "computation_status_detail": "ComputationStatusDetail 2", "firmware_version": "FirmwareVersion 2", "target_firmware_version": "TargetFirmwareVersion 2", "impact_type": "ImpactType 2", "object_type": "firmware.BaseImpact", "model": "Model 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareComponentMeta(t *testing.T) {
	p := []models.FirmwareComponentMeta{}
	var d = &schema.ResourceData{}
	c := `{"Model":"Model %d","Description":"Description %d","IsOobSupported":true,"ComponentLabel":"ComponentLabel %d","PackedVersion":"PackedVersion %d","ClassId":"firmware.ComponentMeta","Vendor":"Vendor %d","Disruption":"Disruption %d","RedfishUrl":"RedfishUrl %d","ComponentType":"ComponentType %d","ImagePath":"ImagePath %d","ObjectType":"firmware.ComponentMeta"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareComponentMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareComponentMeta{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareComponentMeta(p, d)
	expectedOp := []map[string]interface{}{{"model": "Model 1", "description": "Description 1", "is_oob_supported": true, "component_label": "ComponentLabel 1", "packed_version": "PackedVersion 1", "class_id": "firmware.ComponentMeta", "vendor": "Vendor 1", "disruption": "Disruption 1", "redfish_url": "RedfishUrl 1", "component_type": "ComponentType 1", "image_path": "ImagePath 1", "object_type": "firmware.ComponentMeta"}, {"model": "Model 2", "description": "Description 2", "is_oob_supported": true, "component_label": "ComponentLabel 2", "packed_version": "PackedVersion 2", "class_id": "firmware.ComponentMeta", "vendor": "Vendor 2", "disruption": "Disruption 2", "redfish_url": "RedfishUrl 2", "component_type": "ComponentType 2", "image_path": "ImagePath 2", "object_type": "firmware.ComponentMeta"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareDistributableMetaRelationship(t *testing.T) {
	p := []models.FirmwareDistributableMetaRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareDistributableMetaRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareDistributableMetaRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareDistributableMetaRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareFirmwareInventory(t *testing.T) {
	p := []models.FirmwareFirmwareInventory{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"firmware.FirmwareInventory","ObjectType":"firmware.FirmwareInventory","Category":"Category %d","Label":"Label %d","UpdateUri":"UpdateUri %d","Version":"Version %d","Model":"Model %d","Vendor":"Vendor %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareFirmwareInventory(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareFirmwareInventory{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareFirmwareInventory(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "firmware.FirmwareInventory", "object_type": "firmware.FirmwareInventory", "category": "Category 1", "label": "Label 1", "update_uri": "UpdateUri 1", "nr_version": "Version 1", "model": "Model 1", "vendor": "Vendor 1"}, {"class_id": "firmware.FirmwareInventory", "object_type": "firmware.FirmwareInventory", "category": "Category 2", "label": "Label 2", "update_uri": "UpdateUri 2", "nr_version": "Version 2", "model": "Model 2", "vendor": "Vendor 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareRunningFirmwareRelationship(t *testing.T) {
	p := []models.FirmwareRunningFirmwareRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareRunningFirmwareRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareRunningFirmwareRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareRunningFirmwareRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListForecastDefinitionRelationship(t *testing.T) {
	p := []models.ForecastDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListForecastDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ForecastDefinitionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListForecastDefinitionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListGraphicsCardRelationship(t *testing.T) {
	p := []models.GraphicsCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListGraphicsCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.GraphicsCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListGraphicsCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListGraphicsControllerRelationship(t *testing.T) {
	p := []models.GraphicsControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListGraphicsControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.GraphicsControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListGraphicsControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclConstraint(t *testing.T) {
	p := []models.HclConstraint{}
	var d = &schema.ResourceData{}
	c := `{"ConstraintName":"ConstraintName %d","ConstraintValue":"ConstraintValue %d","ClassId":"hcl.Constraint","ObjectType":"hcl.Constraint"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclConstraint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclConstraint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclConstraint(p, d)
	expectedOp := []map[string]interface{}{{"constraint_name": "ConstraintName 1", "constraint_value": "ConstraintValue 1", "class_id": "hcl.Constraint", "object_type": "hcl.Constraint"}, {"constraint_name": "ConstraintName 2", "constraint_value": "ConstraintValue 2", "class_id": "hcl.Constraint", "object_type": "hcl.Constraint"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclHyperflexSoftwareCompatibilityInfoRelationship(t *testing.T) {
	p := []models.HclHyperflexSoftwareCompatibilityInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclHyperflexSoftwareCompatibilityInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclOperatingSystemRelationship(t *testing.T) {
	p := []models.HclOperatingSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclOperatingSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclOperatingSystemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclOperatingSystemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexAlarmRelationship(t *testing.T) {
	p := []models.HyperflexAlarmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexAlarmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexAlarmRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexAlarmRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexBaseClusterRelationship(t *testing.T) {
	p := []models.HyperflexBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexBaseClusterRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexBaseClusterRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexCapabilityInfoRelationship(t *testing.T) {
	p := []models.HyperflexCapabilityInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexCapabilityInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexCapabilityInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexCapabilityInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexClusterProfileRelationship(t *testing.T) {
	p := []models.HyperflexClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexClusterProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexClusterProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexConfigResultEntryRelationship(t *testing.T) {
	p := []models.HyperflexConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexFeatureLimitEntry(t *testing.T) {
	p := []models.HyperflexFeatureLimitEntry{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.FeatureLimitEntry","ClassId":"hyperflex.FeatureLimitEntry","Name":"Name %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexFeatureLimitEntry(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexFeatureLimitEntry{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexFeatureLimitEntry(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "hyperflex.FeatureLimitEntry", "class_id": "hyperflex.FeatureLimitEntry", "name": "Name 1", "value": "Value 1"}, {"object_type": "hyperflex.FeatureLimitEntry", "class_id": "hyperflex.FeatureLimitEntry", "name": "Name 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHealthCheckScriptInfo(t *testing.T) {
	p := []models.HyperflexHealthCheckScriptInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.HealthCheckScriptInfo","HyperflexVersion":"HyperflexVersion %d","ScriptExecuteLocation":"ScriptExecuteLocation %d","ScriptName":"ScriptName %d","Version":"Version %d","ClassId":"hyperflex.HealthCheckScriptInfo","ScriptInput":"ScriptInput %d","AggregateScriptName":"AggregateScriptName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHealthCheckScriptInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHealthCheckScriptInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHealthCheckScriptInfo(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "hyperflex.HealthCheckScriptInfo", "hyperflex_version": "HyperflexVersion 1", "script_execute_location": "ScriptExecuteLocation 1", "script_name": "ScriptName 1", "nr_version": "Version 1", "class_id": "hyperflex.HealthCheckScriptInfo", "script_input": "ScriptInput 1", "aggregate_script_name": "AggregateScriptName 1"}, {"object_type": "hyperflex.HealthCheckScriptInfo", "hyperflex_version": "HyperflexVersion 2", "script_execute_location": "ScriptExecuteLocation 2", "script_name": "ScriptName 2", "nr_version": "Version 2", "class_id": "hyperflex.HealthCheckScriptInfo", "script_input": "ScriptInput 2", "aggregate_script_name": "AggregateScriptName 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxHostMountStatusDt(t *testing.T) {
	p := []models.HyperflexHxHostMountStatusDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxHostMountStatusDt","ObjectType":"hyperflex.HxHostMountStatusDt","Accessibility":"Accessibility %d","HostName":"HostName %d","Mounted":true,"Reason":"Reason %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxHostMountStatusDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxHostMountStatusDt{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxHostMountStatusDt(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.HxHostMountStatusDt", "object_type": "hyperflex.HxHostMountStatusDt", "accessibility": "Accessibility 1", "host_name": "HostName 1", "mounted": true, "reason": "Reason 1"}, {"class_id": "hyperflex.HxHostMountStatusDt", "object_type": "hyperflex.HxHostMountStatusDt", "accessibility": "Accessibility 2", "host_name": "HostName 2", "mounted": true, "reason": "Reason 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxZoneResiliencyInfoDt(t *testing.T) {
	p := []models.HyperflexHxZoneResiliencyInfoDt{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","ClassId":"hyperflex.HxZoneResiliencyInfoDt","ObjectType":"hyperflex.HxZoneResiliencyInfoDt"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxZoneResiliencyInfoDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxZoneResiliencyInfoDt{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxZoneResiliencyInfoDt(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "class_id": "hyperflex.HxZoneResiliencyInfoDt", "object_type": "hyperflex.HxZoneResiliencyInfoDt"}, {"name": "Name 2", "class_id": "hyperflex.HxZoneResiliencyInfoDt", "object_type": "hyperflex.HxZoneResiliencyInfoDt"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxapHostRelationship(t *testing.T) {
	p := []models.HyperflexHxapHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxapHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxapHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxapHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxapHostInterfaceRelationship(t *testing.T) {
	p := []models.HyperflexHxapHostInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxapHostInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxapHostInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxapHostInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxapHostVswitchRelationship(t *testing.T) {
	p := []models.HyperflexHxapHostVswitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxapHostVswitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxapHostVswitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxapHostVswitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxdpVersionRelationship(t *testing.T) {
	p := []models.HyperflexHxdpVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxdpVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxdpVersionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxdpVersionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexIpAddrRange(t *testing.T) {
	p := []models.HyperflexIpAddrRange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.IpAddrRange","ObjectType":"hyperflex.IpAddrRange","EndAddr":"EndAddr %d","Gateway":"Gateway %d","Netmask":"Netmask %d","StartAddr":"StartAddr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexIpAddrRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexIpAddrRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexIpAddrRange(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange", "end_addr": "EndAddr 1", "gateway": "Gateway 1", "netmask": "Netmask 1", "start_addr": "StartAddr 1"}, {"class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange", "end_addr": "EndAddr 2", "gateway": "Gateway 2", "netmask": "Netmask 2", "start_addr": "StartAddr 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexMapClusterIdToProtectionInfo(t *testing.T) {
	p := []models.HyperflexMapClusterIdToProtectionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.MapClusterIdToProtectionInfo","ObjectType":"hyperflex.MapClusterIdToProtectionInfo","ClusterId":"ClusterId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexMapClusterIdToProtectionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexMapClusterIdToProtectionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexMapClusterIdToProtectionInfo(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.MapClusterIdToProtectionInfo", "object_type": "hyperflex.MapClusterIdToProtectionInfo", "cluster_id": "ClusterId 1"}, {"class_id": "hyperflex.MapClusterIdToProtectionInfo", "object_type": "hyperflex.MapClusterIdToProtectionInfo", "cluster_id": "ClusterId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexMapClusterIdToStSnapshotPoint(t *testing.T) {
	p := []models.HyperflexMapClusterIdToStSnapshotPoint{}
	var d = &schema.ResourceData{}
	c := `{"ClusterId":"ClusterId %d","ClassId":"hyperflex.MapClusterIdToStSnapshotPoint","ObjectType":"hyperflex.MapClusterIdToStSnapshotPoint"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexMapClusterIdToStSnapshotPoint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexMapClusterIdToStSnapshotPoint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexMapClusterIdToStSnapshotPoint(p, d)
	expectedOp := []map[string]interface{}{{"cluster_id": "ClusterId 1", "class_id": "hyperflex.MapClusterIdToStSnapshotPoint", "object_type": "hyperflex.MapClusterIdToStSnapshotPoint"}, {"cluster_id": "ClusterId 2", "class_id": "hyperflex.MapClusterIdToStSnapshotPoint", "object_type": "hyperflex.MapClusterIdToStSnapshotPoint"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNamedVlan(t *testing.T) {
	p := []models.HyperflexNamedVlan{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.NamedVlan","ObjectType":"hyperflex.NamedVlan","Name":"Name %d","VlanId":32}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNamedVlan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNamedVlan{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNamedVlan(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "name": "Name 1", "vlan_id": 32}, {"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "name": "Name 2", "vlan_id": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNetworkPort(t *testing.T) {
	p := []models.HyperflexNetworkPort{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.NetworkPort","Name":"Name %d","PortType":"PortType %d","Vlans":"Vlans %d","ClassId":"hyperflex.NetworkPort"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNetworkPort(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNetworkPort{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNetworkPort(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "hyperflex.NetworkPort", "name": "Name 1", "port_type": "PortType 1", "vlans": "Vlans 1", "class_id": "hyperflex.NetworkPort"}, {"object_type": "hyperflex.NetworkPort", "name": "Name 2", "port_type": "PortType 2", "vlans": "Vlans 2", "class_id": "hyperflex.NetworkPort"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNodeRelationship(t *testing.T) {
	p := []models.HyperflexNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNodeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNodeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNodeProfileRelationship(t *testing.T) {
	p := []models.HyperflexNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNodeProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNodeProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexReplicationClusterReferenceToSchedule(t *testing.T) {
	p := []models.HyperflexReplicationClusterReferenceToSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.ReplicationClusterReferenceToSchedule","ObjectType":"hyperflex.ReplicationClusterReferenceToSchedule"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexReplicationClusterReferenceToSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexReplicationClusterReferenceToSchedule{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexReplicationClusterReferenceToSchedule(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.ReplicationClusterReferenceToSchedule", "object_type": "hyperflex.ReplicationClusterReferenceToSchedule"}, {"class_id": "hyperflex.ReplicationClusterReferenceToSchedule", "object_type": "hyperflex.ReplicationClusterReferenceToSchedule"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerFirmwareVersionEntryRelationship(t *testing.T) {
	p := []models.HyperflexServerFirmwareVersionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerFirmwareVersionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerFirmwareVersionEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerFirmwareVersionEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerFirmwareVersionInfo(t *testing.T) {
	p := []models.HyperflexServerFirmwareVersionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ServerPlatform":"ServerPlatform %d","Version":"Version %d","ClassId":"hyperflex.ServerFirmwareVersionInfo","ObjectType":"hyperflex.ServerFirmwareVersionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerFirmwareVersionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerFirmwareVersionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerFirmwareVersionInfo(p, d)
	expectedOp := []map[string]interface{}{{"server_platform": "ServerPlatform 1", "nr_version": "Version 1", "class_id": "hyperflex.ServerFirmwareVersionInfo", "object_type": "hyperflex.ServerFirmwareVersionInfo"}, {"server_platform": "ServerPlatform 2", "nr_version": "Version 2", "class_id": "hyperflex.ServerFirmwareVersionInfo", "object_type": "hyperflex.ServerFirmwareVersionInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerModelEntry(t *testing.T) {
	p := []models.HyperflexServerModelEntry{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Value":"Value %d","ObjectType":"hyperflex.ServerModelEntry","ClassId":"hyperflex.ServerModelEntry"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerModelEntry(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerModelEntry{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerModelEntry(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "value": "Value 1", "object_type": "hyperflex.ServerModelEntry", "class_id": "hyperflex.ServerModelEntry"}, {"name": "Name 2", "value": "Value 2", "object_type": "hyperflex.ServerModelEntry", "class_id": "hyperflex.ServerModelEntry"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionComponentRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionComponentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionComponentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionComponentRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionComponentRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionEntryRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionVersionRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionVersionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionVersionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexVmDisk(t *testing.T) {
	p := []models.HyperflexVmDisk{}
	var d = &schema.ResourceData{}
	c := `{"Bus":"Bus %d","ObjectType":"hyperflex.VmDisk","Name":"Name %d","VirtualDiskReference":"VirtualDiskReference %d","ClassId":"hyperflex.VmDisk","Type":"Type %d","BootOrder":32}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexVmDisk(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexVmDisk{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexVmDisk(p, d)
	expectedOp := []map[string]interface{}{{"bus": "Bus 1", "object_type": "hyperflex.VmDisk", "name": "Name 1", "virtual_disk_reference": "VirtualDiskReference 1", "class_id": "hyperflex.VmDisk", "type": "Type 1", "boot_order": 32}, {"bus": "Bus 2", "object_type": "hyperflex.VmDisk", "name": "Name 2", "virtual_disk_reference": "VirtualDiskReference 2", "class_id": "hyperflex.VmDisk", "type": "Type 2", "boot_order": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexVmInterface(t *testing.T) {
	p := []models.HyperflexVmInterface{}
	var d = &schema.ResourceData{}
	c := `{"Bridge":"Bridge %d","ClassId":"hyperflex.VmInterface","ObjectType":"hyperflex.VmInterface","MacAddress":"MacAddress %d","Model":"Model %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexVmInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexVmInterface{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexVmInterface(p, d)
	expectedOp := []map[string]interface{}{{"bridge": "Bridge 1", "class_id": "hyperflex.VmInterface", "object_type": "hyperflex.VmInterface", "mac_address": "MacAddress 1", "model": "Model 1", "name": "Name 1"}, {"bridge": "Bridge 2", "class_id": "hyperflex.VmInterface", "object_type": "hyperflex.VmInterface", "mac_address": "MacAddress 2", "model": "Model 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexVolumeRelationship(t *testing.T) {
	p := []models.HyperflexVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasConnectorPackRelationship(t *testing.T) {
	p := []models.IaasConnectorPackRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasConnectorPackRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasConnectorPackRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasConnectorPackRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasDeviceStatusRelationship(t *testing.T) {
	p := []models.IaasDeviceStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasDeviceStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasDeviceStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasDeviceStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasLicenseKeysInfo(t *testing.T) {
	p := []models.IaasLicenseKeysInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iaas.LicenseKeysInfo","ObjectType":"iaas.LicenseKeysInfo","Pid":"Pid %d","Count":32,"ExpirationDate":"ExpirationDate %d","LicenseId":"LicenseId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasLicenseKeysInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasLicenseKeysInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasLicenseKeysInfo(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iaas.LicenseKeysInfo", "object_type": "iaas.LicenseKeysInfo", "pid": "Pid 1", "nr_count": 32, "expiration_date": "ExpirationDate 1", "license_id": "LicenseId 1"}, {"class_id": "iaas.LicenseKeysInfo", "object_type": "iaas.LicenseKeysInfo", "pid": "Pid 2", "nr_count": 32, "expiration_date": "ExpirationDate 2", "license_id": "LicenseId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasLicenseUtilizationInfo(t *testing.T) {
	p := []models.IaasLicenseUtilizationInfo{}
	var d = &schema.ResourceData{}
	c := `{"Sku":"Sku %d","ClassId":"iaas.LicenseUtilizationInfo","ObjectType":"iaas.LicenseUtilizationInfo","ActualUsed":32,"Label":"Label %d","LicensedLimit":"LicensedLimit %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasLicenseUtilizationInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasLicenseUtilizationInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasLicenseUtilizationInfo(p, d)
	expectedOp := []map[string]interface{}{{"sku": "Sku 1", "class_id": "iaas.LicenseUtilizationInfo", "object_type": "iaas.LicenseUtilizationInfo", "actual_used": 32, "label": "Label 1", "licensed_limit": "LicensedLimit 1"}, {"sku": "Sku 2", "class_id": "iaas.LicenseUtilizationInfo", "object_type": "iaas.LicenseUtilizationInfo", "actual_used": 32, "label": "Label 2", "licensed_limit": "LicensedLimit 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasMostRunTasksRelationship(t *testing.T) {
	p := []models.IaasMostRunTasksRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasMostRunTasksRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasMostRunTasksRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasMostRunTasksRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasWorkflowSteps(t *testing.T) {
	p := []models.IaasWorkflowSteps{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Status":"Status %d","StatusMessage":"StatusMessage %d","CompletedTime":"CompletedTime %d","ClassId":"iaas.WorkflowSteps","ObjectType":"iaas.WorkflowSteps"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasWorkflowSteps(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasWorkflowSteps{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasWorkflowSteps(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "status": "Status 1", "status_message": "StatusMessage 1", "completed_time": "CompletedTime 1", "class_id": "iaas.WorkflowSteps", "object_type": "iaas.WorkflowSteps"}, {"name": "Name 2", "status": "Status 2", "status_message": "StatusMessage 2", "completed_time": "CompletedTime 2", "class_id": "iaas.WorkflowSteps", "object_type": "iaas.WorkflowSteps"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamAccountPermissions(t *testing.T) {
	p := []models.IamAccountPermissions{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iam.AccountPermissions","ObjectType":"iam.AccountPermissions","AccountName":"AccountName %d","AccountStatus":"AccountStatus %d","AccountIdentifier":"AccountIdentifier %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamAccountPermissions(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamAccountPermissions{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamAccountPermissions(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iam.AccountPermissions", "object_type": "iam.AccountPermissions", "account_name": "AccountName 1", "account_status": "AccountStatus 1", "account_identifier": "AccountIdentifier 1"}, {"class_id": "iam.AccountPermissions", "object_type": "iam.AccountPermissions", "account_name": "AccountName 2", "account_status": "AccountStatus 2", "account_identifier": "AccountIdentifier 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamApiKeyRelationship(t *testing.T) {
	p := []models.IamApiKeyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamApiKeyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamApiKeyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamApiKeyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamAppRegistrationRelationship(t *testing.T) {
	p := []models.IamAppRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamAppRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamAppRegistrationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamAppRegistrationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamDomainGroupRelationship(t *testing.T) {
	p := []models.IamDomainGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamDomainGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamDomainGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamDomainGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointPrivilegeRelationship(t *testing.T) {
	p := []models.IamEndPointPrivilegeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointPrivilegeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointPrivilegeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointPrivilegeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointRoleRelationship(t *testing.T) {
	p := []models.IamEndPointRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointUserRoleRelationship(t *testing.T) {
	p := []models.IamEndPointUserRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointUserRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointUserRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointUserRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamFeatureDefinition(t *testing.T) {
	p := []models.IamFeatureDefinition{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iam.FeatureDefinition","ObjectType":"iam.FeatureDefinition","Feature":"Feature %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamFeatureDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamFeatureDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamFeatureDefinition(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iam.FeatureDefinition", "object_type": "iam.FeatureDefinition", "feature": "Feature 1"}, {"class_id": "iam.FeatureDefinition", "object_type": "iam.FeatureDefinition", "feature": "Feature 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamGroupPermissionToRoles(t *testing.T) {
	p := []models.IamGroupPermissionToRoles{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"iam.GroupPermissionToRoles","ClassId":"iam.GroupPermissionToRoles"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamGroupPermissionToRoles(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamGroupPermissionToRoles{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamGroupPermissionToRoles(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "iam.GroupPermissionToRoles", "class_id": "iam.GroupPermissionToRoles"}, {"object_type": "iam.GroupPermissionToRoles", "class_id": "iam.GroupPermissionToRoles"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIdpRelationship(t *testing.T) {
	p := []models.IamIdpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIdpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIdpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIdpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIdpReferenceRelationship(t *testing.T) {
	p := []models.IamIdpReferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIdpReferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIdpReferenceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIdpReferenceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIpAddressRelationship(t *testing.T) {
	p := []models.IamIpAddressRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIpAddressRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIpAddressRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIpAddressRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamLdapGroupRelationship(t *testing.T) {
	p := []models.IamLdapGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamLdapGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamLdapGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamLdapGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamLdapProviderRelationship(t *testing.T) {
	p := []models.IamLdapProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamLdapProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamLdapProviderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamLdapProviderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamOAuthTokenRelationship(t *testing.T) {
	p := []models.IamOAuthTokenRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamOAuthTokenRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamOAuthTokenRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamOAuthTokenRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPermissionRelationship(t *testing.T) {
	p := []models.IamPermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPermissionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPermissionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPermissionToRoles(t *testing.T) {
	p := []models.IamPermissionToRoles{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iam.PermissionToRoles","ObjectType":"iam.PermissionToRoles"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPermissionToRoles(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPermissionToRoles{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPermissionToRoles(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iam.PermissionToRoles", "object_type": "iam.PermissionToRoles"}, {"class_id": "iam.PermissionToRoles", "object_type": "iam.PermissionToRoles"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPrivilegeRelationship(t *testing.T) {
	p := []models.IamPrivilegeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPrivilegeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPrivilegeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPrivilegeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPrivilegeSetRelationship(t *testing.T) {
	p := []models.IamPrivilegeSetRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPrivilegeSetRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPrivilegeSetRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPrivilegeSetRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamResourcePermissionRelationship(t *testing.T) {
	p := []models.IamResourcePermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamResourcePermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamResourcePermissionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamResourcePermissionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamResourceRolesRelationship(t *testing.T) {
	p := []models.IamResourceRolesRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamResourceRolesRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamResourceRolesRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamResourceRolesRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamRoleRelationship(t *testing.T) {
	p := []models.IamRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamSessionRelationship(t *testing.T) {
	p := []models.IamSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamSessionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamSessionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserRelationship(t *testing.T) {
	p := []models.IamUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserGroupRelationship(t *testing.T) {
	p := []models.IamUserGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserPreferenceRelationship(t *testing.T) {
	p := []models.IamUserPreferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserPreferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserPreferenceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserPreferenceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInfraMetaData(t *testing.T) {
	p := []models.InfraMetaData{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"infra.MetaData","Name":"Name %d","Value":"Value %d","ClassId":"infra.MetaData"}`

	//test when the response is empty
	ffOpEmpty := flattenListInfraMetaData(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InfraMetaData{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInfraMetaData(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "infra.MetaData", "name": "Name 1", "value": "Value 1", "class_id": "infra.MetaData"}, {"object_type": "infra.MetaData", "name": "Name 2", "value": "Value 2", "class_id": "infra.MetaData"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInventoryGenericInventoryRelationship(t *testing.T) {
	p := []models.InventoryGenericInventoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListInventoryGenericInventoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InventoryGenericInventoryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInventoryGenericInventoryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInventoryGenericInventoryHolderRelationship(t *testing.T) {
	p := []models.InventoryGenericInventoryHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListInventoryGenericInventoryHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InventoryGenericInventoryHolderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInventoryGenericInventoryHolderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpLeaseRelationship(t *testing.T) {
	p := []models.IppoolIpLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpLeaseRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpLeaseRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpV4Block(t *testing.T) {
	p := []models.IppoolIpV4Block{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"ippool.IpV4Block","ObjectType":"ippool.IpV4Block","Size":32,"From":"From %d","To":"To %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpV4Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpV4Block{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpV4Block(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "ippool.IpV4Block", "object_type": "ippool.IpV4Block", "size": 32, "from": "From 1", "to": "To 1"}, {"class_id": "ippool.IpV4Block", "object_type": "ippool.IpV4Block", "size": 32, "from": "From 2", "to": "To 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpV6Block(t *testing.T) {
	p := []models.IppoolIpV6Block{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ippool.IpV6Block","Size":32,"To":"To %d","From":"From %d","ClassId":"ippool.IpV6Block"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpV6Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpV6Block{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpV6Block(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "ippool.IpV6Block", "size": 32, "to": "To 1", "from": "From 1", "class_id": "ippool.IpV6Block"}, {"object_type": "ippool.IpV6Block", "size": 32, "to": "To 2", "from": "From 2", "class_id": "ippool.IpV6Block"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolPoolRelationship(t *testing.T) {
	p := []models.IppoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolPoolRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolPoolRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolShadowBlockRelationship(t *testing.T) {
	p := []models.IppoolShadowBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolShadowBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolShadowBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolShadowBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolShadowPoolRelationship(t *testing.T) {
	p := []models.IppoolShadowPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolShadowPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolShadowPoolRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolShadowPoolRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIqnpoolBlockRelationship(t *testing.T) {
	p := []models.IqnpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIqnpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IqnpoolBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIqnpoolBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIqnpoolIqnSuffixBlock(t *testing.T) {
	p := []models.IqnpoolIqnSuffixBlock{}
	var d = &schema.ResourceData{}
	c := `{"To":32,"From":32,"Suffix":"Suffix %d","ClassId":"iqnpool.IqnSuffixBlock","ObjectType":"iqnpool.IqnSuffixBlock","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenListIqnpoolIqnSuffixBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IqnpoolIqnSuffixBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIqnpoolIqnSuffixBlock(p, d)
	expectedOp := []map[string]interface{}{{"to": 32, "from": 32, "suffix": "Suffix 1", "class_id": "iqnpool.IqnSuffixBlock", "object_type": "iqnpool.IqnSuffixBlock", "size": 32}, {"to": 32, "from": 32, "suffix": "Suffix 2", "class_id": "iqnpool.IqnSuffixBlock", "object_type": "iqnpool.IqnSuffixBlock", "size": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAciCniProfileRelationship(t *testing.T) {
	p := []models.KubernetesAciCniProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAciCniProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAciCniProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAciCniProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAciCniTenantClusterAllocationRelationship(t *testing.T) {
	p := []models.KubernetesAciCniTenantClusterAllocationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAciCniTenantClusterAllocationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAciCniTenantClusterAllocationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAciCniTenantClusterAllocationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAddon(t *testing.T) {
	p := []models.KubernetesAddon{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.Addon","Name":"Name %d","ClassId":"kubernetes.Addon"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAddon(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAddon{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAddon(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "kubernetes.Addon", "name": "Name 1", "class_id": "kubernetes.Addon"}, {"object_type": "kubernetes.Addon", "name": "Name 2", "class_id": "kubernetes.Addon"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesClusterProfileRelationship(t *testing.T) {
	p := []models.KubernetesClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesClusterProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesClusterProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesConfigResultEntryRelationship(t *testing.T) {
	p := []models.KubernetesConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesEssentialAddon(t *testing.T) {
	p := []models.KubernetesEssentialAddon{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.EssentialAddon","ClassId":"kubernetes.EssentialAddon","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesEssentialAddon(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesEssentialAddon{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesEssentialAddon(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "kubernetes.EssentialAddon", "class_id": "kubernetes.EssentialAddon", "name": "Name 1"}, {"object_type": "kubernetes.EssentialAddon", "class_id": "kubernetes.EssentialAddon", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeAddress(t *testing.T) {
	p := []models.KubernetesNodeAddress{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeAddress","ObjectType":"kubernetes.NodeAddress","Address":"Address %d","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeAddress(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeAddress{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeAddress(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.NodeAddress", "object_type": "kubernetes.NodeAddress", "address": "Address 1", "type": "Type 1"}, {"class_id": "kubernetes.NodeAddress", "object_type": "kubernetes.NodeAddress", "address": "Address 2", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupLabel(t *testing.T) {
	p := []models.KubernetesNodeGroupLabel{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeGroupLabel","ObjectType":"kubernetes.NodeGroupLabel","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupLabel(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupLabel{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupLabel(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.NodeGroupLabel", "object_type": "kubernetes.NodeGroupLabel", "key": "Key 1", "value": "Value 1"}, {"class_id": "kubernetes.NodeGroupLabel", "object_type": "kubernetes.NodeGroupLabel", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupProfileRelationship(t *testing.T) {
	p := []models.KubernetesNodeGroupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupTaint(t *testing.T) {
	p := []models.KubernetesNodeGroupTaint{}
	var d = &schema.ResourceData{}
	c := `{"Effect":"Effect %d","ClassId":"kubernetes.NodeGroupTaint","ObjectType":"kubernetes.NodeGroupTaint","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupTaint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupTaint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupTaint(p, d)
	expectedOp := []map[string]interface{}{{"effect": "Effect 1", "class_id": "kubernetes.NodeGroupTaint", "object_type": "kubernetes.NodeGroupTaint", "key": "Key 1", "value": "Value 1"}, {"effect": "Effect 2", "class_id": "kubernetes.NodeGroupTaint", "object_type": "kubernetes.NodeGroupTaint", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeProfileRelationship(t *testing.T) {
	p := []models.KubernetesNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeStatus(t *testing.T) {
	p := []models.KubernetesNodeStatus{}
	var d = &schema.ResourceData{}
	c := `{"Type":"Type %d","ClassId":"kubernetes.NodeStatus","ObjectType":"kubernetes.NodeStatus","Status":"Status %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeStatus(p, d)
	expectedOp := []map[string]interface{}{{"type": "Type 1", "class_id": "kubernetes.NodeStatus", "object_type": "kubernetes.NodeStatus", "status": "Status 1"}, {"type": "Type 2", "class_id": "kubernetes.NodeStatus", "object_type": "kubernetes.NodeStatus", "status": "Status 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesVirtualMachineInfrastructureProviderRelationship(t *testing.T) {
	p := []models.KubernetesVirtualMachineInfrastructureProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesVirtualMachineInfrastructureProviderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListLicenseLicenseInfoRelationship(t *testing.T) {
	p := []models.LicenseLicenseInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListLicenseLicenseInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.LicenseLicenseInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListLicenseLicenseInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMacpoolBlock(t *testing.T) {
	p := []models.MacpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"macpool.Block","From":"From %d","To":"To %d","ObjectType":"macpool.Block","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenListMacpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MacpoolBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMacpoolBlock(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "macpool.Block", "from": "From 1", "to": "To 1", "object_type": "macpool.Block", "size": 32}, {"class_id": "macpool.Block", "from": "From 2", "to": "To 2", "object_type": "macpool.Block", "size": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMacpoolIdBlockRelationship(t *testing.T) {
	p := []models.MacpoolIdBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMacpoolIdBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MacpoolIdBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMacpoolIdBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListManagementInterfaceRelationship(t *testing.T) {
	p := []models.ManagementInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListManagementInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ManagementInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListManagementInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryArrayRelationship(t *testing.T) {
	p := []models.MemoryArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryArrayRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryArrayRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryGoal(t *testing.T) {
	p := []models.MemoryPersistentMemoryGoal{}
	var d = &schema.ResourceData{}
	c := `{"SocketId":"SocketId %d","MemoryModePercentage":32,"PersistentMemoryType":"PersistentMemoryType %d","ClassId":"memory.PersistentMemoryGoal","ObjectType":"memory.PersistentMemoryGoal"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryGoal(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryGoal{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryGoal(p, d)
	expectedOp := []map[string]interface{}{{"socket_id": "SocketId 1", "memory_mode_percentage": 32, "persistent_memory_type": "PersistentMemoryType 1", "class_id": "memory.PersistentMemoryGoal", "object_type": "memory.PersistentMemoryGoal"}, {"socket_id": "SocketId 2", "memory_mode_percentage": 32, "persistent_memory_type": "PersistentMemoryType 2", "class_id": "memory.PersistentMemoryGoal", "object_type": "memory.PersistentMemoryGoal"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryLogicalNamespace(t *testing.T) {
	p := []models.MemoryPersistentMemoryLogicalNamespace{}
	var d = &schema.ResourceData{}
	c := `{"SocketId":32,"SocketMemoryId":"SocketMemoryId %d","ClassId":"memory.PersistentMemoryLogicalNamespace","ObjectType":"memory.PersistentMemoryLogicalNamespace","Capacity":32,"Mode":"Mode %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryLogicalNamespace(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryLogicalNamespace{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryLogicalNamespace(p, d)
	expectedOp := []map[string]interface{}{{"socket_id": 32, "socket_memory_id": "SocketMemoryId 1", "class_id": "memory.PersistentMemoryLogicalNamespace", "object_type": "memory.PersistentMemoryLogicalNamespace", "capacity": 32, "mode": "Mode 1", "name": "Name 1"}, {"socket_id": 32, "socket_memory_id": "SocketMemoryId 2", "class_id": "memory.PersistentMemoryLogicalNamespace", "object_type": "memory.PersistentMemoryLogicalNamespace", "capacity": 32, "mode": "Mode 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryNamespaceRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryNamespaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryNamespaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryNamespaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryNamespaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryNamespaceConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryNamespaceConfigResultRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryRegionRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryRegionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryRegionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryRegionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryRegionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryUnitRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryUnitRelationship(t *testing.T) {
	p := []models.MemoryUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaAccessPrivilege(t *testing.T) {
	p := []models.MetaAccessPrivilege{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"meta.AccessPrivilege","ObjectType":"meta.AccessPrivilege","Method":"Method %d","Privilege":"Privilege %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaAccessPrivilege(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaAccessPrivilege{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaAccessPrivilege(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "meta.AccessPrivilege", "object_type": "meta.AccessPrivilege", "method": "Method 1", "privilege": "Privilege 1"}, {"class_id": "meta.AccessPrivilege", "object_type": "meta.AccessPrivilege", "method": "Method 2", "privilege": "Privilege 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaDisplayNameDefinition(t *testing.T) {
	p := []models.MetaDisplayNameDefinition{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Format":"Format %d","ClassId":"meta.DisplayNameDefinition","ObjectType":"meta.DisplayNameDefinition","IncludeAncestor":true}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaDisplayNameDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaDisplayNameDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaDisplayNameDefinition(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "format": "Format 1", "class_id": "meta.DisplayNameDefinition", "object_type": "meta.DisplayNameDefinition", "include_ancestor": true}, {"name": "Name 2", "format": "Format 2", "class_id": "meta.DisplayNameDefinition", "object_type": "meta.DisplayNameDefinition", "include_ancestor": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaPropDefinition(t *testing.T) {
	p := []models.MetaPropDefinition{}
	var d = &schema.ResourceData{}
	c := `{"SearchWeight":32.000000,"ApiAccess":"ApiAccess %d","Name":"Name %d","OpSecurity":"OpSecurity %d","ClassId":"meta.PropDefinition","ObjectType":"meta.PropDefinition"}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaPropDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaPropDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaPropDefinition(p, d)
	expectedOp := []map[string]interface{}{{"search_weight": 32.000000, "api_access": "ApiAccess 1", "name": "Name 1", "op_security": "OpSecurity 1", "class_id": "meta.PropDefinition", "object_type": "meta.PropDefinition"}, {"search_weight": 32.000000, "api_access": "ApiAccess 2", "name": "Name 2", "op_security": "OpSecurity 2", "class_id": "meta.PropDefinition", "object_type": "meta.PropDefinition"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaRelationshipDefinition(t *testing.T) {
	p := []models.MetaRelationshipDefinition{}
	var d = &schema.ResourceData{}
	c := `{"ApiAccess":"ApiAccess %d","Export":true,"Type":"Type %d","PeerRelName":"PeerRelName %d","ExportWithPeer":true,"ClassId":"meta.RelationshipDefinition","ObjectType":"meta.RelationshipDefinition","Name":"Name %d","Collection":true,"PeerSync":true}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaRelationshipDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaRelationshipDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaRelationshipDefinition(p, d)
	expectedOp := []map[string]interface{}{{"api_access": "ApiAccess 1", "export": true, "type": "Type 1", "peer_rel_name": "PeerRelName 1", "export_with_peer": true, "class_id": "meta.RelationshipDefinition", "object_type": "meta.RelationshipDefinition", "name": "Name 1", "collection": true, "peer_sync": true}, {"api_access": "ApiAccess 2", "export": true, "type": "Type 2", "peer_rel_name": "PeerRelName 2", "export_with_peer": true, "class_id": "meta.RelationshipDefinition", "object_type": "meta.RelationshipDefinition", "name": "Name 2", "collection": true, "peer_sync": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMoBaseMoRelationship(t *testing.T) {
	p := []models.MoBaseMoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMoBaseMoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MoBaseMoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMoBaseMoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMoTag(t *testing.T) {
	p := []models.MoTag{}
	var d = &schema.ResourceData{}
	c := `{"Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMoTag(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MoTag{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMoTag(p, d)
	expectedOp := []map[string]interface{}{{"key": "Key 1", "value": "Value 1"}, {"key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNetworkElementRelationship(t *testing.T) {
	p := []models.NetworkElementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNetworkElementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NetworkElementRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNetworkElementRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiaapiDetail(t *testing.T) {
	p := []models.NiaapiDetail{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","ClassId":"niaapi.Detail","ObjectType":"niaapi.Detail","Chksum":"Chksum %d","Filename":"Filename %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiaapiDetail(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiaapiDetail{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiaapiDetail(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "class_id": "niaapi.Detail", "object_type": "niaapi.Detail", "chksum": "Chksum 1", "filename": "Filename 1"}, {"name": "Name 2", "class_id": "niaapi.Detail", "object_type": "niaapi.Detail", "chksum": "Chksum 2", "filename": "Filename 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiaapiRevisionInfo(t *testing.T) {
	p := []models.NiaapiRevisionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niaapi.RevisionInfo","ObjectType":"niaapi.RevisionInfo","RevisionComment":"RevisionComment %d","RevisionNo":"RevisionNo %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiaapiRevisionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiaapiRevisionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiaapiRevisionInfo(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "niaapi.RevisionInfo", "object_type": "niaapi.RevisionInfo", "revision_comment": "RevisionComment 1", "revision_no": "RevisionNo 1"}, {"class_id": "niaapi.RevisionInfo", "object_type": "niaapi.RevisionInfo", "revision_comment": "RevisionComment 2", "revision_no": "RevisionNo 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiatelemetryLogicalLink(t *testing.T) {
	p := []models.NiatelemetryLogicalLink{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.LogicalLink","LinkAddr1":"LinkAddr1 %d","LinkAddr2":"LinkAddr2 %d","DbId":32,"IsPresent":true,"LinkState":"LinkState %d","ObjectType":"niatelemetry.LogicalLink","Uptime":"Uptime %d","LinkType":"LinkType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiatelemetryLogicalLink(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiatelemetryLogicalLink{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiatelemetryLogicalLink(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "niatelemetry.LogicalLink", "link_addr1": "LinkAddr1 1", "link_addr2": "LinkAddr2 1", "db_id": 32, "is_present": true, "link_state": "LinkState 1", "object_type": "niatelemetry.LogicalLink", "uptime": "Uptime 1", "link_type": "LinkType 1"}, {"class_id": "niatelemetry.LogicalLink", "link_addr1": "LinkAddr1 2", "link_addr2": "LinkAddr2 2", "db_id": 32, "is_present": true, "link_state": "LinkState 2", "object_type": "niatelemetry.LogicalLink", "uptime": "Uptime 2", "link_type": "LinkType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNotificationAbstractCondition(t *testing.T) {
	p := []models.NotificationAbstractCondition{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"notification.AbstractCondition","ClassId":"notification.AbstractCondition"}`

	//test when the response is empty
	ffOpEmpty := flattenListNotificationAbstractCondition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NotificationAbstractCondition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNotificationAbstractCondition(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "notification.AbstractCondition", "class_id": "notification.AbstractCondition"}, {"object_type": "notification.AbstractCondition", "class_id": "notification.AbstractCondition"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNotificationAction(t *testing.T) {
	p := []models.NotificationAction{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"notification.Action","ObjectType":"notification.Action"}`

	//test when the response is empty
	ffOpEmpty := flattenListNotificationAction(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NotificationAction{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNotificationAction(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "notification.Action", "object_type": "notification.Action"}, {"class_id": "notification.Action", "object_type": "notification.Action"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNtpAuthNtpServer(t *testing.T) {
	p := []models.NtpAuthNtpServer{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"ntp.AuthNtpServer","KeyType":"KeyType %d","ServerName":"ServerName %d","SymKeyId":32,"SymKeyValue":"SymKeyValue %d","ObjectType":"ntp.AuthNtpServer"}`

	//test when the response is empty
	ffOpEmpty := flattenListNtpAuthNtpServer(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NtpAuthNtpServer{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNtpAuthNtpServer(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "ntp.AuthNtpServer", "key_type": "KeyType 1", "server_name": "ServerName 1", "sym_key_id": 32, "sym_key_value": "SymKeyValue 1", "object_type": "ntp.AuthNtpServer"}, {"class_id": "ntp.AuthNtpServer", "key_type": "KeyType 2", "server_name": "ServerName 2", "sym_key_id": 32, "sym_key_value": "SymKeyValue 2", "object_type": "ntp.AuthNtpServer"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremImagePackage(t *testing.T) {
	p := []models.OnpremImagePackage{}
	var d = &schema.ResourceData{}
	c := `{"FilePath":"FilePath %d","FileSha":"FileSha %d","FileSize":32,"Version":"Version %d","Name":"Name %d","ClassId":"onprem.ImagePackage","ObjectType":"onprem.ImagePackage","PackageType":"PackageType %d","Filename":"Filename %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremImagePackage(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremImagePackage{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremImagePackage(p, d)
	expectedOp := []map[string]interface{}{{"file_path": "FilePath 1", "file_sha": "FileSha 1", "file_size": 32, "nr_version": "Version 1", "name": "Name 1", "class_id": "onprem.ImagePackage", "object_type": "onprem.ImagePackage", "package_type": "PackageType 1", "filename": "Filename 1"}, {"file_path": "FilePath 2", "file_sha": "FileSha 2", "file_size": 32, "nr_version": "Version 2", "name": "Name 2", "class_id": "onprem.ImagePackage", "object_type": "onprem.ImagePackage", "package_type": "PackageType 2", "filename": "Filename 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremUpgradeNote(t *testing.T) {
	p := []models.OnpremUpgradeNote{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.UpgradeNote","ObjectType":"onprem.UpgradeNote","Message":"Message %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremUpgradeNote(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremUpgradeNote{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremUpgradeNote(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "onprem.UpgradeNote", "object_type": "onprem.UpgradeNote", "message": "Message 1"}, {"class_id": "onprem.UpgradeNote", "object_type": "onprem.UpgradeNote", "message": "Message 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremUpgradePhase(t *testing.T) {
	p := []models.OnpremUpgradePhase{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","ElapsedTime":32,"Failed":true,"ClassId":"onprem.UpgradePhase","ObjectType":"onprem.UpgradePhase","Message":"Message %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremUpgradePhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremUpgradePhase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremUpgradePhase(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "elapsed_time": 32, "failed": true, "class_id": "onprem.UpgradePhase", "object_type": "onprem.UpgradePhase", "message": "Message 1"}, {"name": "Name 2", "elapsed_time": 32, "failed": true, "class_id": "onprem.UpgradePhase", "object_type": "onprem.UpgradePhase", "message": "Message 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOprsKvpair(t *testing.T) {
	p := []models.OprsKvpair{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"oprs.Kvpair","ObjectType":"oprs.Kvpair","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOprsKvpair(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OprsKvpair{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOprsKvpair(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "oprs.Kvpair", "object_type": "oprs.Kvpair", "key": "Key 1", "value": "Value 1"}, {"class_id": "oprs.Kvpair", "object_type": "oprs.Kvpair", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOrganizationOrganizationRelationship(t *testing.T) {
	p := []models.OrganizationOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOrganizationOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OrganizationOrganizationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOrganizationOrganizationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsConfigurationFileRelationship(t *testing.T) {
	p := []models.OsConfigurationFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsConfigurationFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsConfigurationFileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsConfigurationFileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsDistributionRelationship(t *testing.T) {
	p := []models.OsDistributionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsDistributionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsDistributionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsDistributionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsPlaceHolder(t *testing.T) {
	p := []models.OsPlaceHolder{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.PlaceHolder","ObjectType":"os.PlaceHolder","IsValueSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenListOsPlaceHolder(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsPlaceHolder{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsPlaceHolder(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "os.PlaceHolder", "object_type": "os.PlaceHolder", "is_value_set": true}, {"class_id": "os.PlaceHolder", "object_type": "os.PlaceHolder", "is_value_set": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsServerConfig(t *testing.T) {
	p := []models.OsServerConfig{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"os.ServerConfig","ClassId":"os.ServerConfig","InstallTarget":"InstallTarget %d","SerialNumber":"SerialNumber %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsServerConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsServerConfig{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsServerConfig(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "os.ServerConfig", "class_id": "os.ServerConfig", "install_target": "InstallTarget 1", "serial_number": "SerialNumber 1"}, {"object_type": "os.ServerConfig", "class_id": "os.ServerConfig", "install_target": "InstallTarget 2", "serial_number": "SerialNumber 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsValidationInformation(t *testing.T) {
	p := []models.OsValidationInformation{}
	var d = &schema.ResourceData{}
	c := `{"Status":"Status %d","StepName":"StepName %d","ErrorMsg":"ErrorMsg %d","ObjectType":"os.ValidationInformation","ClassId":"os.ValidationInformation"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsValidationInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsValidationInformation{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsValidationInformation(p, d)
	expectedOp := []map[string]interface{}{{"status": "Status 1", "step_name": "StepName 1", "error_msg": "ErrorMsg 1", "object_type": "os.ValidationInformation", "class_id": "os.ValidationInformation"}, {"status": "Status 2", "step_name": "StepName 2", "error_msg": "ErrorMsg 2", "object_type": "os.ValidationInformation", "class_id": "os.ValidationInformation"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciCoprocessorCardRelationship(t *testing.T) {
	p := []models.PciCoprocessorCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciCoprocessorCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciCoprocessorCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciCoprocessorCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciDeviceRelationship(t *testing.T) {
	p := []models.PciDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciLinkRelationship(t *testing.T) {
	p := []models.PciLinkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciLinkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciLinkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciLinkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciSwitchRelationship(t *testing.T) {
	p := []models.PciSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciSwitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciSwitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyAbstractConfigProfileRelationship(t *testing.T) {
	p := []models.PolicyAbstractConfigProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyAbstractConfigProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyAbstractConfigProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyAbstractConfigProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyAbstractPolicyRelationship(t *testing.T) {
	p := []models.PolicyAbstractPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyAbstractPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyAbstractPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyAbstractPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyinventoryJobInfo(t *testing.T) {
	p := []models.PolicyinventoryJobInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"policyinventory.JobInfo","ExecutionStatus":"ExecutionStatus %d","PolicyId":"PolicyId %d","PolicyName":"PolicyName %d","ClassId":"policyinventory.JobInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyinventoryJobInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyinventoryJobInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyinventoryJobInfo(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "policyinventory.JobInfo", "execution_status": "ExecutionStatus 1", "policy_id": "PolicyId 1", "policy_name": "PolicyName 1", "class_id": "policyinventory.JobInfo"}, {"object_type": "policyinventory.JobInfo", "execution_status": "ExecutionStatus 2", "policy_id": "PolicyId 2", "policy_name": "PolicyName 2", "class_id": "policyinventory.JobInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortGroupRelationship(t *testing.T) {
	p := []models.PortGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortMacBindingRelationship(t *testing.T) {
	p := []models.PortMacBindingRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortMacBindingRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortMacBindingRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortMacBindingRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortSubGroupRelationship(t *testing.T) {
	p := []models.PortSubGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortSubGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortSubGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortSubGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListProcessorUnitRelationship(t *testing.T) {
	p := []models.ProcessorUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListProcessorUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ProcessorUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListProcessorUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecommendationPhysicalItemRelationship(t *testing.T) {
	p := []models.RecommendationPhysicalItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecommendationPhysicalItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecommendationPhysicalItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecommendationPhysicalItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecoveryBackupProfileRelationship(t *testing.T) {
	p := []models.RecoveryBackupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecoveryBackupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecoveryBackupProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecoveryBackupProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecoveryConfigResultEntryRelationship(t *testing.T) {
	p := []models.RecoveryConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecoveryConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecoveryConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecoveryConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourceGroupRelationship(t *testing.T) {
	p := []models.ResourceGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourceGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourceGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourceGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourcePerTypeCombinedSelector(t *testing.T) {
	p := []models.ResourcePerTypeCombinedSelector{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"resource.PerTypeCombinedSelector","CombinedSelector":"CombinedSelector %d","EmptyFilter":true,"SelectorObjectType":"SelectorObjectType %d","ClassId":"resource.PerTypeCombinedSelector"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourcePerTypeCombinedSelector(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourcePerTypeCombinedSelector{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourcePerTypeCombinedSelector(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "resource.PerTypeCombinedSelector", "combined_selector": "CombinedSelector 1", "empty_filter": true, "selector_object_type": "SelectorObjectType 1", "class_id": "resource.PerTypeCombinedSelector"}, {"object_type": "resource.PerTypeCombinedSelector", "combined_selector": "CombinedSelector 2", "empty_filter": true, "selector_object_type": "SelectorObjectType 2", "class_id": "resource.PerTypeCombinedSelector"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourceSelector(t *testing.T) {
	p := []models.ResourceSelector{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"resource.Selector","Selector":"Selector %d","ClassId":"resource.Selector"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourceSelector(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourceSelector{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourceSelector(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "resource.Selector", "selector": "Selector 1", "class_id": "resource.Selector"}, {"object_type": "resource.Selector", "selector": "Selector 2", "class_id": "resource.Selector"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdcardPartition(t *testing.T) {
	p := []models.SdcardPartition{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"sdcard.Partition","ObjectType":"sdcard.Partition","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdcardPartition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdcardPartition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdcardPartition(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "sdcard.Partition", "object_type": "sdcard.Partition", "type": "Type 1"}, {"class_id": "sdcard.Partition", "object_type": "sdcard.Partition", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanNetworkConfigurationType(t *testing.T) {
	p := []models.SdwanNetworkConfigurationType{}
	var d = &schema.ResourceData{}
	c := `{"PortGroup":"PortGroup %d","Vlan":32,"ClassId":"sdwan.NetworkConfigurationType","ObjectType":"sdwan.NetworkConfigurationType","NetworkType":"NetworkType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanNetworkConfigurationType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanNetworkConfigurationType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanNetworkConfigurationType(p, d)
	expectedOp := []map[string]interface{}{{"port_group": "PortGroup 1", "vlan": 32, "class_id": "sdwan.NetworkConfigurationType", "object_type": "sdwan.NetworkConfigurationType", "network_type": "NetworkType 1"}, {"port_group": "PortGroup 2", "vlan": 32, "class_id": "sdwan.NetworkConfigurationType", "object_type": "sdwan.NetworkConfigurationType", "network_type": "NetworkType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanProfileRelationship(t *testing.T) {
	p := []models.SdwanProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanRouterNodeRelationship(t *testing.T) {
	p := []models.SdwanRouterNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanRouterNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanRouterNodeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanRouterNodeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanTemplateInputsType(t *testing.T) {
	p := []models.SdwanTemplateInputsType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"sdwan.TemplateInputsType","Required":true,"Editable":true,"Key":"Key %d","ObjectType":"sdwan.TemplateInputsType","Type":"Type %d","Title":"Title %d","Value":"Value %d","Template":"Template %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanTemplateInputsType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanTemplateInputsType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanTemplateInputsType(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "sdwan.TemplateInputsType", "required": true, "editable": true, "key": "Key 1", "object_type": "sdwan.TemplateInputsType", "type": "Type 1", "title": "Title 1", "value": "Value 1", "template": "Template 1"}, {"class_id": "sdwan.TemplateInputsType", "required": true, "editable": true, "key": "Key 2", "object_type": "sdwan.TemplateInputsType", "type": "Type 2", "title": "Title 2", "value": "Value 2", "template": "Template 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSecurityUnitRelationship(t *testing.T) {
	p := []models.SecurityUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSecurityUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SecurityUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSecurityUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListServerConfigChangeDetailRelationship(t *testing.T) {
	p := []models.ServerConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListServerConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ServerConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListServerConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListServerConfigResultEntryRelationship(t *testing.T) {
	p := []models.ServerConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListServerConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ServerConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListServerConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSnmpTrap(t *testing.T) {
	p := []models.SnmpTrap{}
	var d = &schema.ResourceData{}
	c := `{"User":"User %d","Enabled":true,"ObjectType":"snmp.Trap","Version":"Version %d","Destination":"Destination %d","Type":"Type %d","Community":"Community %d","ClassId":"snmp.Trap","Port":32}`

	//test when the response is empty
	ffOpEmpty := flattenListSnmpTrap(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SnmpTrap{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSnmpTrap(p, d)
	expectedOp := []map[string]interface{}{{"user": "User 1", "enabled": true, "object_type": "snmp.Trap", "nr_version": "Version 1", "destination": "Destination 1", "type": "Type 1", "community": "Community 1", "class_id": "snmp.Trap", "port": 32}, {"user": "User 2", "enabled": true, "object_type": "snmp.Trap", "nr_version": "Version 2", "destination": "Destination 2", "type": "Type 2", "community": "Community 2", "class_id": "snmp.Trap", "port": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSnmpUser(t *testing.T) {
	p := []models.SnmpUser{}
	var d = &schema.ResourceData{}
	c := `{"IsAuthPasswordSet":true,"Name":"Name %d","AuthType":"AuthType %d","PrivacyType":"PrivacyType %d","SecurityLevel":"SecurityLevel %d","ClassId":"snmp.User","ObjectType":"snmp.User","IsPrivacyPasswordSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenListSnmpUser(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SnmpUser{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSnmpUser(p, d)
	expectedOp := []map[string]interface{}{{"is_auth_password_set": true, "name": "Name 1", "auth_type": "AuthType 1", "privacy_type": "PrivacyType 1", "security_level": "SecurityLevel 1", "class_id": "snmp.User", "object_type": "snmp.User", "is_privacy_password_set": true}, {"is_auth_password_set": true, "name": "Name 2", "auth_type": "AuthType 2", "privacy_type": "PrivacyType 2", "security_level": "SecurityLevel 2", "class_id": "snmp.User", "object_type": "snmp.User", "is_privacy_password_set": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwareHyperflexDistributableRelationship(t *testing.T) {
	p := []models.SoftwareHyperflexDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwareHyperflexDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwareHyperflexDistributableRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwareHyperflexDistributableRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwareUcsdDistributableRelationship(t *testing.T) {
	p := []models.SoftwareUcsdDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwareUcsdDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwareUcsdDistributableRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwareUcsdDistributableRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwarerepositoryConstraintModels(t *testing.T) {
	p := []models.SoftwarerepositoryConstraintModels{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"softwarerepository.ConstraintModels","ObjectType":"softwarerepository.ConstraintModels","MinVersion":"MinVersion %d","Name":"Name %d","PlatformRegex":"PlatformRegex %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwarerepositoryConstraintModels(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwarerepositoryConstraintModels{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwarerepositoryConstraintModels(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "softwarerepository.ConstraintModels", "object_type": "softwarerepository.ConstraintModels", "min_version": "MinVersion 1", "name": "Name 1", "platform_regex": "PlatformRegex 1"}, {"class_id": "softwarerepository.ConstraintModels", "object_type": "softwarerepository.ConstraintModels", "min_version": "MinVersion 2", "name": "Name 2", "platform_regex": "PlatformRegex 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageBaseInitiator(t *testing.T) {
	p := []models.StorageBaseInitiator{}
	var d = &schema.ResourceData{}
	c := `{"Iqn":"Iqn %d","Name":"Name %d","Type":"Type %d","Wwn":"Wwn %d","ClassId":"storage.BaseInitiator","ObjectType":"storage.BaseInitiator"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageBaseInitiator(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageBaseInitiator{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageBaseInitiator(p, d)
	expectedOp := []map[string]interface{}{{"iqn": "Iqn 1", "name": "Name 1", "type": "Type 1", "wwn": "Wwn 1", "class_id": "storage.BaseInitiator", "object_type": "storage.BaseInitiator"}, {"iqn": "Iqn 2", "name": "Name 2", "type": "Type 2", "wwn": "Wwn 2", "class_id": "storage.BaseInitiator", "object_type": "storage.BaseInitiator"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageControllerRelationship(t *testing.T) {
	p := []models.StorageControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageDiskGroupRelationship(t *testing.T) {
	p := []models.StorageDiskGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageDiskGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageDiskGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageDiskGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageDiskGroupPolicyRelationship(t *testing.T) {
	p := []models.StorageDiskGroupPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageDiskGroupPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageDiskGroupPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageDiskGroupPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureRelationship(t *testing.T) {
	p := []models.StorageEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureDiskRelationship(t *testing.T) {
	p := []models.StorageEnclosureDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureDiskRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureDiskRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureDiskSlotEpRelationship(t *testing.T) {
	p := []models.StorageEnclosureDiskSlotEpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureDiskSlotEpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureDiskSlotEpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureDiskSlotEpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashControllerRelationship(t *testing.T) {
	p := []models.StorageFlexFlashControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashControllerPropsRelationship(t *testing.T) {
	p := []models.StorageFlexFlashControllerPropsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashControllerPropsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashControllerPropsRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashControllerPropsRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashPhysicalDriveRelationship(t *testing.T) {
	p := []models.StorageFlexFlashPhysicalDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashPhysicalDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashPhysicalDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashPhysicalDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageFlexFlashVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilControllerRelationship(t *testing.T) {
	p := []models.StorageFlexUtilControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilPhysicalDriveRelationship(t *testing.T) {
	p := []models.StorageFlexUtilPhysicalDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilPhysicalDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilPhysicalDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilPhysicalDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageFlexUtilVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHitachiParityGroupRelationship(t *testing.T) {
	p := []models.StorageHitachiParityGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHitachiParityGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHitachiParityGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHitachiParityGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHyperFlexStorageContainerRelationship(t *testing.T) {
	p := []models.StorageHyperFlexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHyperFlexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHyperFlexStorageContainerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHyperFlexStorageContainerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHyperFlexVolumeRelationship(t *testing.T) {
	p := []models.StorageHyperFlexVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHyperFlexVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHyperFlexVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHyperFlexVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageItemRelationship(t *testing.T) {
	p := []models.StorageItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageLocalDisk(t *testing.T) {
	p := []models.StorageLocalDisk{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"storage.LocalDisk","ObjectType":"storage.LocalDisk","SlotNumber":32}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageLocalDisk(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageLocalDisk{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageLocalDisk(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "storage.LocalDisk", "object_type": "storage.LocalDisk", "slot_number": 32}, {"class_id": "storage.LocalDisk", "object_type": "storage.LocalDisk", "slot_number": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppAggregateRelationship(t *testing.T) {
	p := []models.StorageNetAppAggregateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppAggregateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppAggregateRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppAggregateRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppExportPolicyRule(t *testing.T) {
	p := []models.StorageNetAppExportPolicyRule{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"storage.NetAppExportPolicyRule","Index":32,"User":"User %d","ClassId":"storage.NetAppExportPolicyRule"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppExportPolicyRule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppExportPolicyRule{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppExportPolicyRule(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "storage.NetAppExportPolicyRule", "index": 32, "user": "User 1", "class_id": "storage.NetAppExportPolicyRule"}, {"object_type": "storage.NetAppExportPolicyRule", "index": 32, "user": "User 2", "class_id": "storage.NetAppExportPolicyRule"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppInitiatorGroupRelationship(t *testing.T) {
	p := []models.StorageNetAppInitiatorGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppInitiatorGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppInitiatorGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppInitiatorGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskExtensionRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskExtensionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskExtensionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskUsageRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskUsageRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskUsageRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskUsageRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskUsageRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureHostRelationship(t *testing.T) {
	p := []models.StoragePureHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureHostGroupRelationship(t *testing.T) {
	p := []models.StoragePureHostGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureHostGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureHostGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureHostGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureReplicationBlackout(t *testing.T) {
	p := []models.StoragePureReplicationBlackout{}
	var d = &schema.ResourceData{}
	c := `{"Start":"Start %d","End":"End %d","ClassId":"storage.PureReplicationBlackout","ObjectType":"storage.PureReplicationBlackout"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureReplicationBlackout(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureReplicationBlackout{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureReplicationBlackout(p, d)
	expectedOp := []map[string]interface{}{{"start": "Start 1", "end": "End 1", "class_id": "storage.PureReplicationBlackout", "object_type": "storage.PureReplicationBlackout"}, {"start": "Start 2", "end": "End 2", "class_id": "storage.PureReplicationBlackout", "object_type": "storage.PureReplicationBlackout"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureVolumeRelationship(t *testing.T) {
	p := []models.StoragePureVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSasExpanderRelationship(t *testing.T) {
	p := []models.StorageSasExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSasExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSasExpanderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSasExpanderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSasPortRelationship(t *testing.T) {
	p := []models.StorageSasPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSasPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSasPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSasPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSpanRelationship(t *testing.T) {
	p := []models.StorageSpanRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSpanRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSpanRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSpanRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSpanGroup(t *testing.T) {
	p := []models.StorageSpanGroup{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"storage.SpanGroup","ObjectType":"storage.SpanGroup"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSpanGroup(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSpanGroup{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSpanGroup(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "storage.SpanGroup", "object_type": "storage.SpanGroup"}, {"class_id": "storage.SpanGroup", "object_type": "storage.SpanGroup"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageStoragePolicyRelationship(t *testing.T) {
	p := []models.StorageStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageStoragePolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageStoragePolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVdMemberEpRelationship(t *testing.T) {
	p := []models.StorageVdMemberEpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVdMemberEpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVdMemberEpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVdMemberEpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveConfig(t *testing.T) {
	p := []models.StorageVirtualDriveConfig{}
	var d = &schema.ResourceData{}
	c := `{"StripSize":"StripSize %d","Size":32,"AccessPolicy":"AccessPolicy %d","DiskGroupPolicy":"DiskGroupPolicy %d","ObjectType":"storage.VirtualDriveConfig","DriveCache":"DriveCache %d","Vdid":"Vdid %d","WritePolicy":"WritePolicy %d","DiskGroupName":"DiskGroupName %d","ClassId":"storage.VirtualDriveConfig","IoPolicy":"IoPolicy %d","ReadPolicy":"ReadPolicy %d","BootDrive":true,"Name":"Name %d","ExpandToAvailable":true}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveConfig{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveConfig(p, d)
	expectedOp := []map[string]interface{}{{"strip_size": "StripSize 1", "size": 32, "access_policy": "AccessPolicy 1", "disk_group_policy": "DiskGroupPolicy 1", "object_type": "storage.VirtualDriveConfig", "drive_cache": "DriveCache 1", "vdid": "Vdid 1", "write_policy": "WritePolicy 1", "disk_group_name": "DiskGroupName 1", "class_id": "storage.VirtualDriveConfig", "io_policy": "IoPolicy 1", "read_policy": "ReadPolicy 1", "boot_drive": true, "name": "Name 1", "expand_to_available": true}, {"strip_size": "StripSize 2", "size": 32, "access_policy": "AccessPolicy 2", "disk_group_policy": "DiskGroupPolicy 2", "object_type": "storage.VirtualDriveConfig", "drive_cache": "DriveCache 2", "vdid": "Vdid 2", "write_policy": "WritePolicy 2", "disk_group_name": "DiskGroupName 2", "class_id": "storage.VirtualDriveConfig", "io_policy": "IoPolicy 2", "read_policy": "ReadPolicy 2", "boot_drive": true, "name": "Name 2", "expand_to_available": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveContainerRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveContainerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveContainerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveExtensionRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveExtensionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveExtensionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSyslogLocalClientBase(t *testing.T) {
	p := []models.SyslogLocalClientBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"syslog.LocalClientBase","ObjectType":"syslog.LocalClientBase","MinSeverity":"MinSeverity %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSyslogLocalClientBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SyslogLocalClientBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSyslogLocalClientBase(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "syslog.LocalClientBase", "object_type": "syslog.LocalClientBase", "min_severity": "MinSeverity 1"}, {"class_id": "syslog.LocalClientBase", "object_type": "syslog.LocalClientBase", "min_severity": "MinSeverity 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSyslogRemoteClientBase(t *testing.T) {
	p := []models.SyslogRemoteClientBase{}
	var d = &schema.ResourceData{}
	c := `{"Hostname":"Hostname %d","MinSeverity":"MinSeverity %d","Port":32,"Protocol":"Protocol %d","ClassId":"syslog.RemoteClientBase","ObjectType":"syslog.RemoteClientBase","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenListSyslogRemoteClientBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SyslogRemoteClientBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSyslogRemoteClientBase(p, d)
	expectedOp := []map[string]interface{}{{"hostname": "Hostname 1", "min_severity": "MinSeverity 1", "port": 32, "protocol": "Protocol 1", "class_id": "syslog.RemoteClientBase", "object_type": "syslog.RemoteClientBase", "enabled": true}, {"hostname": "Hostname 2", "min_severity": "MinSeverity 2", "port": 32, "protocol": "Protocol 2", "class_id": "syslog.RemoteClientBase", "object_type": "syslog.RemoteClientBase", "enabled": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamAction(t *testing.T) {
	p := []models.TamAction{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Type":"Type %d","AffectedObjectType":"AffectedObjectType %d","AlertType":"AlertType %d","ClassId":"tam.Action","OperationType":"OperationType %d","ObjectType":"tam.Action"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamAction(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamAction{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamAction(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "type": "Type 1", "affected_object_type": "AffectedObjectType 1", "alert_type": "AlertType 1", "class_id": "tam.Action", "operation_type": "OperationType 1", "object_type": "tam.Action"}, {"name": "Name 2", "type": "Type 2", "affected_object_type": "AffectedObjectType 2", "alert_type": "AlertType 2", "class_id": "tam.Action", "operation_type": "OperationType 2", "object_type": "tam.Action"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamApiDataSource(t *testing.T) {
	p := []models.TamApiDataSource{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"tam.ApiDataSource","ObjectType":"tam.ApiDataSource","Type":"Type %d","Name":"Name %d","MoType":"MoType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamApiDataSource(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamApiDataSource{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamApiDataSource(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "tam.ApiDataSource", "object_type": "tam.ApiDataSource", "type": "Type 1", "name": "Name 1", "mo_type": "MoType 1"}, {"class_id": "tam.ApiDataSource", "object_type": "tam.ApiDataSource", "type": "Type 2", "name": "Name 2", "mo_type": "MoType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamS3DataSource(t *testing.T) {
	p := []models.TamS3DataSource{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Type":"Type %d","ClassId":"tam.S3DataSource","ObjectType":"tam.S3DataSource","S3Path":"S3Path %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamS3DataSource(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamS3DataSource{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamS3DataSource(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "type": "Type 1", "class_id": "tam.S3DataSource", "object_type": "tam.S3DataSource", "s3_path": "S3Path 1"}, {"name": "Name 2", "type": "Type 2", "class_id": "tam.S3DataSource", "object_type": "tam.S3DataSource", "s3_path": "S3Path 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUcsdConnectorPack(t *testing.T) {
	p := []models.UcsdConnectorPack{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ucsd.ConnectorPack","Version":"Version %d","ConnectorFeature":"ConnectorFeature %d","DownloadedVersion":"DownloadedVersion %d","Name":"Name %d","ClassId":"ucsd.ConnectorPack","State":"State %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListUcsdConnectorPack(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UcsdConnectorPack{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUcsdConnectorPack(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "ucsd.ConnectorPack", "nr_version": "Version 1", "connector_feature": "ConnectorFeature 1", "downloaded_version": "DownloadedVersion 1", "name": "Name 1", "class_id": "ucsd.ConnectorPack", "state": "State 1"}, {"object_type": "ucsd.ConnectorPack", "nr_version": "Version 2", "connector_feature": "ConnectorFeature 2", "downloaded_version": "DownloadedVersion 2", "name": "Name 2", "class_id": "ucsd.ConnectorPack", "state": "State 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUuidpoolBlockRelationship(t *testing.T) {
	p := []models.UuidpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListUuidpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UuidpoolBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUuidpoolBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUuidpoolUuidBlock(t *testing.T) {
	p := []models.UuidpoolUuidBlock{}
	var d = &schema.ResourceData{}
	c := `{"Size":32,"ClassId":"uuidpool.UuidBlock","ObjectType":"uuidpool.UuidBlock","To":"To %d","From":"From %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListUuidpoolUuidBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UuidpoolUuidBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUuidpoolUuidBlock(p, d)
	expectedOp := []map[string]interface{}{{"size": 32, "class_id": "uuidpool.UuidBlock", "object_type": "uuidpool.UuidBlock", "to": "To 1", "from": "From 1"}, {"size": 32, "class_id": "uuidpool.UuidBlock", "object_type": "uuidpool.UuidBlock", "to": "To 2", "from": "From 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationBaseNetworkRelationship(t *testing.T) {
	p := []models.VirtualizationBaseNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationBaseNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationBaseNetworkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationBaseNetworkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationNetworkInterface(t *testing.T) {
	p := []models.VirtualizationNetworkInterface{}
	var d = &schema.ResourceData{}
	c := `{"AdaptorType":"AdaptorType %d","Bridge":"Bridge %d","ConnectAtPowerOn":true,"MacAddress":"MacAddress %d","Name":"Name %d","ObjectType":"virtualization.NetworkInterface","ClassId":"virtualization.NetworkInterface","DirectPathIo":true}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationNetworkInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationNetworkInterface{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationNetworkInterface(p, d)
	expectedOp := []map[string]interface{}{{"adaptor_type": "AdaptorType 1", "bridge": "Bridge 1", "connect_at_power_on": true, "mac_address": "MacAddress 1", "name": "Name 1", "object_type": "virtualization.NetworkInterface", "class_id": "virtualization.NetworkInterface", "direct_path_io": true}, {"adaptor_type": "AdaptorType 2", "bridge": "Bridge 2", "connect_at_power_on": true, "mac_address": "MacAddress 2", "name": "Name 2", "object_type": "virtualization.NetworkInterface", "class_id": "virtualization.NetworkInterface", "direct_path_io": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVirtualMachineDisk(t *testing.T) {
	p := []models.VirtualizationVirtualMachineDisk{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.VirtualMachineDisk","VirtualDiskReference":"VirtualDiskReference %d","Bus":"Bus %d","Order":32,"ClassId":"virtualization.VirtualMachineDisk","Name":"Name %d","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVirtualMachineDisk(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVirtualMachineDisk{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVirtualMachineDisk(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "virtualization.VirtualMachineDisk", "virtual_disk_reference": "VirtualDiskReference 1", "bus": "Bus 1", "order": 32, "class_id": "virtualization.VirtualMachineDisk", "name": "Name 1", "type": "Type 1"}, {"object_type": "virtualization.VirtualMachineDisk", "virtual_disk_reference": "VirtualDiskReference 2", "bus": "Bus 2", "order": 32, "class_id": "virtualization.VirtualMachineDisk", "name": "Name 2", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDatastoreRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDatastoreRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDatastoreRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDatastoreRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDatastoreRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDistributedNetworkRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDistributedNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDistributedNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDistributedNetworkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDistributedNetworkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDistributedSwitchRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDistributedSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDistributedSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDistributedSwitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDistributedSwitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareHostRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareVlanRange(t *testing.T) {
	p := []models.VirtualizationVmwareVlanRange{}
	var d = &schema.ResourceData{}
	c := `{"VlanRangeEnd":32,"VlanRangeStart":32,"ClassId":"virtualization.VmwareVlanRange","ObjectType":"virtualization.VmwareVlanRange"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareVlanRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareVlanRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareVlanRange(p, d)
	expectedOp := []map[string]interface{}{{"vlan_range_end": 32, "vlan_range_start": 32, "class_id": "virtualization.VmwareVlanRange", "object_type": "virtualization.VmwareVlanRange"}, {"vlan_range_end": 32, "vlan_range_start": 32, "class_id": "virtualization.VmwareVlanRange", "object_type": "virtualization.VmwareVlanRange"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVmediaMapping(t *testing.T) {
	p := []models.VmediaMapping{}
	var d = &schema.ResourceData{}
	c := `{"RemotePath":"RemotePath %d","Username":"Username %d","FileLocation":"FileLocation %d","HostName":"HostName %d","AuthenticationProtocol":"AuthenticationProtocol %d","SanitizedFileLocation":"SanitizedFileLocation %d","IsPasswordSet":true,"ClassId":"vmedia.Mapping","DeviceType":"DeviceType %d","ObjectType":"vmedia.Mapping","VolumeName":"VolumeName %d","MountProtocol":"MountProtocol %d","RemoteFile":"RemoteFile %d","MountOptions":"MountOptions %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVmediaMapping(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VmediaMapping{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVmediaMapping(p, d)
	expectedOp := []map[string]interface{}{{"remote_path": "RemotePath 1", "username": "Username 1", "file_location": "FileLocation 1", "host_name": "HostName 1", "authentication_protocol": "AuthenticationProtocol 1", "sanitized_file_location": "SanitizedFileLocation 1", "is_password_set": true, "class_id": "vmedia.Mapping", "device_type": "DeviceType 1", "object_type": "vmedia.Mapping", "volume_name": "VolumeName 1", "mount_protocol": "MountProtocol 1", "remote_file": "RemoteFile 1", "mount_options": "MountOptions 1"}, {"remote_path": "RemotePath 2", "username": "Username 2", "file_location": "FileLocation 2", "host_name": "HostName 2", "authentication_protocol": "AuthenticationProtocol 2", "sanitized_file_location": "SanitizedFileLocation 2", "is_password_set": true, "class_id": "vmedia.Mapping", "device_type": "DeviceType 2", "object_type": "vmedia.Mapping", "volume_name": "VolumeName 2", "mount_protocol": "MountProtocol 2", "remote_file": "RemoteFile 2", "mount_options": "MountOptions 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicEthIfRelationship(t *testing.T) {
	p := []models.VnicEthIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicEthIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicEthIfRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicEthIfRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicEthNetworkPolicyRelationship(t *testing.T) {
	p := []models.VnicEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicEthNetworkPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicEthNetworkPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicFcIfRelationship(t *testing.T) {
	p := []models.VnicFcIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicFcIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicFcIfRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicFcIfRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicVifStatus(t *testing.T) {
	p := []models.VnicVifStatus{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Reason":"Reason %d","Status":"Status %d","ClassId":"vnic.VifStatus","ObjectType":"vnic.VifStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicVifStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicVifStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicVifStatus(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "reason": "Reason 1", "status": "Status 1", "class_id": "vnic.VifStatus", "object_type": "vnic.VifStatus"}, {"name": "Name 2", "reason": "Reason 2", "status": "Status 2", "class_id": "vnic.VifStatus", "object_type": "vnic.VifStatus"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowApi(t *testing.T) {
	p := []models.WorkflowApi{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.Api","SkipOnCondition":"SkipOnCondition %d","AssetTargetMoid":"AssetTargetMoid %d","StartDelay":32,"Body":"Body %d","Timeout":32,"Name":"Name %d","Description":"Description %d","ContentType":"ContentType %d","ClassId":"workflow.Api","ErrorContentType":"ErrorContentType %d","Label":"Label %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowApi(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowApi{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowApi(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.Api", "skip_on_condition": "SkipOnCondition 1", "asset_target_moid": "AssetTargetMoid 1", "start_delay": 32, "body": "Body 1", "timeout": 32, "name": "Name 1", "description": "Description 1", "content_type": "ContentType 1", "class_id": "workflow.Api", "error_content_type": "ErrorContentType 1", "label": "Label 1"}, {"object_type": "workflow.Api", "skip_on_condition": "SkipOnCondition 2", "asset_target_moid": "AssetTargetMoid 2", "start_delay": 32, "body": "Body 2", "timeout": 32, "name": "Name 2", "description": "Description 2", "content_type": "ContentType 2", "class_id": "workflow.Api", "error_content_type": "ErrorContentType 2", "label": "Label 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowBaseDataType(t *testing.T) {
	p := []models.WorkflowBaseDataType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.BaseDataType","ObjectType":"workflow.BaseDataType","Required":true,"Description":"Description %d","Name":"Name %d","Label":"Label %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowBaseDataType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowBaseDataType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowBaseDataType(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.BaseDataType", "object_type": "workflow.BaseDataType", "required": true, "description": "Description 1", "name": "Name 1", "label": "Label 1"}, {"class_id": "workflow.BaseDataType", "object_type": "workflow.BaseDataType", "required": true, "description": "Description 2", "name": "Name 2", "label": "Label 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowDynamicWorkflowActionTaskList(t *testing.T) {
	p := []models.WorkflowDynamicWorkflowActionTaskList{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.DynamicWorkflowActionTaskList","ObjectType":"workflow.DynamicWorkflowActionTaskList","Action":"Action %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowDynamicWorkflowActionTaskList(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowDynamicWorkflowActionTaskList{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowDynamicWorkflowActionTaskList(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.DynamicWorkflowActionTaskList", "object_type": "workflow.DynamicWorkflowActionTaskList", "action": "Action 1"}, {"class_id": "workflow.DynamicWorkflowActionTaskList", "object_type": "workflow.DynamicWorkflowActionTaskList", "action": "Action 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowMessage(t *testing.T) {
	p := []models.WorkflowMessage{}
	var d = &schema.ResourceData{}
	c := `{"Message":"Message %d","Severity":"Severity %d","ClassId":"workflow.Message","ObjectType":"workflow.Message"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowMessage(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowMessage{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowMessage(p, d)
	expectedOp := []map[string]interface{}{{"message": "Message 1", "severity": "Severity 1", "class_id": "workflow.Message", "object_type": "workflow.Message"}, {"message": "Message 2", "severity": "Severity 2", "class_id": "workflow.Message", "object_type": "workflow.Message"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowParameterSet(t *testing.T) {
	p := []models.WorkflowParameterSet{}
	var d = &schema.ResourceData{}
	c := `{"ControlParameter":"ControlParameter %d","Name":"Name %d","Value":"Value %d","ClassId":"workflow.ParameterSet","ObjectType":"workflow.ParameterSet","Condition":"Condition %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowParameterSet(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowParameterSet{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowParameterSet(p, d)
	expectedOp := []map[string]interface{}{{"control_parameter": "ControlParameter 1", "name": "Name 1", "value": "Value 1", "class_id": "workflow.ParameterSet", "object_type": "workflow.ParameterSet", "condition": "Condition 1"}, {"control_parameter": "ControlParameter 2", "name": "Name 2", "value": "Value 2", "class_id": "workflow.ParameterSet", "object_type": "workflow.ParameterSet", "condition": "Condition 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowRollbackTask(t *testing.T) {
	p := []models.WorkflowRollbackTask{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Version":32,"Description":"Description %d","ClassId":"workflow.RollbackTask","ObjectType":"workflow.RollbackTask","TaskMoid":"TaskMoid %d","CatalogMoid":"CatalogMoid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowRollbackTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowRollbackTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowRollbackTask(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "nr_version": 32, "description": "Description 1", "class_id": "workflow.RollbackTask", "object_type": "workflow.RollbackTask", "task_moid": "TaskMoid 1", "catalog_moid": "CatalogMoid 1"}, {"name": "Name 2", "nr_version": 32, "description": "Description 2", "class_id": "workflow.RollbackTask", "object_type": "workflow.RollbackTask", "task_moid": "TaskMoid 2", "catalog_moid": "CatalogMoid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowRollbackWorkflowTask(t *testing.T) {
	p := []models.WorkflowRollbackWorkflowTask{}
	var d = &schema.ResourceData{}
	c := `{"Description":"Description %d","RollbackTaskName":"RollbackTaskName %d","TaskPath":"TaskPath %d","ClassId":"workflow.RollbackWorkflowTask","ObjectType":"workflow.RollbackWorkflowTask","RefName":"RefName %d","Name":"Name %d","RollbackCompleted":true,"Status":"Status %d","TaskInfoMoid":"TaskInfoMoid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowRollbackWorkflowTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowRollbackWorkflowTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowRollbackWorkflowTask(p, d)
	expectedOp := []map[string]interface{}{{"description": "Description 1", "rollback_task_name": "RollbackTaskName 1", "task_path": "TaskPath 1", "class_id": "workflow.RollbackWorkflowTask", "object_type": "workflow.RollbackWorkflowTask", "ref_name": "RefName 1", "name": "Name 1", "rollback_completed": true, "status": "Status 1", "task_info_moid": "TaskInfoMoid 1"}, {"description": "Description 2", "rollback_task_name": "RollbackTaskName 2", "task_path": "TaskPath 2", "class_id": "workflow.RollbackWorkflowTask", "object_type": "workflow.RollbackWorkflowTask", "ref_name": "RefName 2", "name": "Name 2", "rollback_completed": true, "status": "Status 2", "task_info_moid": "TaskInfoMoid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskDefinitionRelationship(t *testing.T) {
	p := []models.WorkflowTaskDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskDefinitionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskDefinitionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskInfoRelationship(t *testing.T) {
	p := []models.WorkflowTaskInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskRetryInfo(t *testing.T) {
	p := []models.WorkflowTaskRetryInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.TaskRetryInfo","Status":"Status %d","TaskInstId":"TaskInstId %d","ClassId":"workflow.TaskRetryInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskRetryInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskRetryInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskRetryInfo(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.TaskRetryInfo", "status": "Status 1", "task_inst_id": "TaskInstId 1", "class_id": "workflow.TaskRetryInfo"}, {"object_type": "workflow.TaskRetryInfo", "status": "Status 2", "task_inst_id": "TaskInstId 2", "class_id": "workflow.TaskRetryInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowUiInputFilter(t *testing.T) {
	p := []models.WorkflowUiInputFilter{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","UserHelpMessage":"UserHelpMessage %d","ClassId":"workflow.UiInputFilter","ObjectType":"workflow.UiInputFilter"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowUiInputFilter(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowUiInputFilter{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowUiInputFilter(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "user_help_message": "UserHelpMessage 1", "class_id": "workflow.UiInputFilter", "object_type": "workflow.UiInputFilter"}, {"name": "Name 2", "user_help_message": "UserHelpMessage 2", "class_id": "workflow.UiInputFilter", "object_type": "workflow.UiInputFilter"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowWorkflowInfoRelationship(t *testing.T) {
	p := []models.WorkflowWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowWorkflowInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowWorkflowInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowWorkflowTask(t *testing.T) {
	p := []models.WorkflowWorkflowTask{}
	var d = &schema.ResourceData{}
	c := `{"Label":"Label %d","Name":"Name %d","Description":"Description %d","ClassId":"workflow.WorkflowTask","ObjectType":"workflow.WorkflowTask"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowWorkflowTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowWorkflowTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowWorkflowTask(p, d)
	expectedOp := []map[string]interface{}{{"label": "Label 1", "name": "Name 1", "description": "Description 1", "class_id": "workflow.WorkflowTask", "object_type": "workflow.WorkflowTask"}, {"label": "Label 2", "name": "Name 2", "description": "Description 2", "class_id": "workflow.WorkflowTask", "object_type": "workflow.WorkflowTask"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListX509Certificate(t *testing.T) {
	p := []models.X509Certificate{}
	var d = &schema.ResourceData{}
	c := `{"PemCertificate":"PemCertificate %d","Sha256Fingerprint":"Sha256Fingerprint %d","SignatureAlgorithm":"SignatureAlgorithm %d","ObjectType":"x509.Certificate","ClassId":"x509.Certificate"}`

	//test when the response is empty
	ffOpEmpty := flattenListX509Certificate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.X509Certificate{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListX509Certificate(p, d)
	expectedOp := []map[string]interface{}{{"pem_certificate": "PemCertificate 1", "sha256_fingerprint": "Sha256Fingerprint 1", "signature_algorithm": "SignatureAlgorithm 1", "object_type": "x509.Certificate", "class_id": "x509.Certificate"}, {"pem_certificate": "PemCertificate 2", "sha256_fingerprint": "Sha256Fingerprint 2", "signature_algorithm": "SignatureAlgorithm 2", "object_type": "x509.Certificate", "class_id": "x509.Certificate"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenMapAccessAddressType(t *testing.T) {
	p := models.AccessAddressType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"access.AddressType","ObjectType":"access.AddressType","EnableIpV4":true,"EnableIpV6":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapAccessAddressType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAccessAddressType(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "access.AddressType", "object_type": "access.AddressType", "enable_ip_v4": true, "enable_ip_v6": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAdapterUnitRelationship(t *testing.T) {
	p := models.AdapterUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAdapterUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAdapterUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAdapterUnitExpanderRelationship(t *testing.T) {
	p := models.AdapterUnitExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAdapterUnitExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAdapterUnitExpanderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceCertRenewalPhase(t *testing.T) {
	p := models.ApplianceCertRenewalPhase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"appliance.CertRenewalPhase","Name":"Name %d","Failed":true,"Message":"Message %d","ObjectType":"appliance.CertRenewalPhase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceCertRenewalPhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceCertRenewalPhase(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "appliance.CertRenewalPhase", "name": "Name 1", "failed": true, "message": "Message 1", "object_type": "appliance.CertRenewalPhase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceDataExportPolicyRelationship(t *testing.T) {
	p := models.ApplianceDataExportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceDataExportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceDataExportPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceGroupStatusRelationship(t *testing.T) {
	p := models.ApplianceGroupStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceGroupStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceGroupStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceImageBundleRelationship(t *testing.T) {
	p := models.ApplianceImageBundleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceImageBundleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceImageBundleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceNodeInfoRelationship(t *testing.T) {
	p := models.ApplianceNodeInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceNodeInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceNodeInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceNodeStatusRelationship(t *testing.T) {
	p := models.ApplianceNodeStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceNodeStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceNodeStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceSystemInfoRelationship(t *testing.T) {
	p := models.ApplianceSystemInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceSystemInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceSystemInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceSystemStatusRelationship(t *testing.T) {
	p := models.ApplianceSystemStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceSystemStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceSystemStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetClaimSignature(t *testing.T) {
	p := models.AssetClaimSignature{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.ClaimSignature","ObjectType":"asset.ClaimSignature"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetClaimSignature(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetClaimSignature(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.ClaimSignature", "object_type": "asset.ClaimSignature"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetClusterMemberRelationship(t *testing.T) {
	p := models.AssetClusterMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetClusterMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetClusterMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetContractInformation(t *testing.T) {
	p := models.AssetContractInformation{}
	var d = &schema.ResourceData{}
	c := `{"ContractNumber":"ContractNumber %d","LineStatus":"LineStatus %d","ClassId":"asset.ContractInformation","ObjectType":"asset.ContractInformation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetContractInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetContractInformation(p, d)[0]
	expectedOp := map[string]interface{}{"contract_number": "ContractNumber 1", "line_status": "LineStatus 1", "class_id": "asset.ContractInformation", "object_type": "asset.ContractInformation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetCustomerInformation(t *testing.T) {
	p := models.AssetCustomerInformation{}
	var d = &schema.ResourceData{}
	c := `{"Id":"Id %d","Name":"Name %d","ClassId":"asset.CustomerInformation","ObjectType":"asset.CustomerInformation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetCustomerInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetCustomerInformation(p, d)[0]
	expectedOp := map[string]interface{}{"id": "Id 1", "name": "Name 1", "class_id": "asset.CustomerInformation", "object_type": "asset.CustomerInformation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentRelationship(t *testing.T) {
	p := models.AssetDeploymentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentDeviceInformation(t *testing.T) {
	p := models.AssetDeploymentDeviceInformation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.DeploymentDeviceInformation","OldDeviceId":"OldDeviceId %d","Description":"Description %d","ObjectType":"asset.DeploymentDeviceInformation","MlbProductName":"MlbProductName %d","ItemType":"ItemType %d","OldInstanceId":"OldInstanceId %d","ApplicationName":"ApplicationName %d","MlbProductId":32,"OldDeviceStatusId":32,"OldDeviceStatusDescription":"OldDeviceStatusDescription %d","InstanceId":"InstanceId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentDeviceInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentDeviceInformation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.DeploymentDeviceInformation", "old_device_id": "OldDeviceId 1", "description": "Description 1", "object_type": "asset.DeploymentDeviceInformation", "mlb_product_name": "MlbProductName 1", "item_type": "ItemType 1", "old_instance_id": "OldInstanceId 1", "application_name": "ApplicationName 1", "mlb_product_id": 32, "old_device_status_id": 32, "old_device_status_description": "OldDeviceStatusDescription 1", "instance_id": "InstanceId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceClaimRelationship(t *testing.T) {
	p := models.AssetDeviceClaimRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceClaimRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceClaimRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceConfigurationRelationship(t *testing.T) {
	p := models.AssetDeviceConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceConnectionRelationship(t *testing.T) {
	p := models.AssetDeviceConnectionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceConnectionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceConnectionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceContractInformationRelationship(t *testing.T) {
	p := models.AssetDeviceContractInformationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceContractInformationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceContractInformationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceInformation(t *testing.T) {
	p := models.AssetDeviceInformation{}
	var d = &schema.ResourceData{}
	c := `{"MlbOfferType":"MlbOfferType %d","OldDeviceStatusDescription":"OldDeviceStatusDescription %d","OldDeviceId":"OldDeviceId %d","ProductType":"ProductType %d","UnitOfMeasure":"UnitOfMeasure %d","InstanceId":"InstanceId %d","MlbProductId":32,"OldDeviceStatusId":32,"OldInstanceId":"OldInstanceId %d","MlbProductName":"MlbProductName %d","ProductFamily":"ProductFamily %d","ClassId":"asset.DeviceInformation","ObjectType":"asset.DeviceInformation","ItemType":"ItemType %d","ApplicationName":"ApplicationName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceInformation(p, d)[0]
	expectedOp := map[string]interface{}{"mlb_offer_type": "MlbOfferType 1", "old_device_status_description": "OldDeviceStatusDescription 1", "old_device_id": "OldDeviceId 1", "product_type": "ProductType 1", "unit_of_measure": "UnitOfMeasure 1", "instance_id": "InstanceId 1", "mlb_product_id": 32, "old_device_status_id": 32, "old_instance_id": "OldInstanceId 1", "mlb_product_name": "MlbProductName 1", "product_family": "ProductFamily 1", "class_id": "asset.DeviceInformation", "object_type": "asset.DeviceInformation", "item_type": "ItemType 1", "application_name": "ApplicationName 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceRegistrationRelationship(t *testing.T) {
	p := models.AssetDeviceRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceRegistrationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceStatistics(t *testing.T) {
	p := models.AssetDeviceStatistics{}
	var d = &schema.ResourceData{}
	c := `{"MembershipRatio":32.000000,"MemoryMirroringFactor":32.000000,"ObjectType":"asset.DeviceStatistics","ClassId":"asset.DeviceStatistics","ClusterName":"ClusterName %d","Connected":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceStatistics(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceStatistics(p, d)[0]
	expectedOp := map[string]interface{}{"membership_ratio": 32.000000, "memory_mirroring_factor": 32.000000, "object_type": "asset.DeviceStatistics", "class_id": "asset.DeviceStatistics", "cluster_name": "ClusterName 1", "connected": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetGlobalUltimate(t *testing.T) {
	p := models.AssetGlobalUltimate{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","ClassId":"asset.GlobalUltimate","ObjectType":"asset.GlobalUltimate","Id":"Id %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetGlobalUltimate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetGlobalUltimate(p, d)[0]
	expectedOp := map[string]interface{}{"name": "Name 1", "class_id": "asset.GlobalUltimate", "object_type": "asset.GlobalUltimate", "id": "Id 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetProductInformation(t *testing.T) {
	p := models.AssetProductInformation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.ProductInformation","SubType":"SubType %d","Number":"Number %d","Group":"Group %d","Description":"Description %d","ObjectType":"asset.ProductInformation","Family":"Family %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetProductInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetProductInformation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.ProductInformation", "sub_type": "SubType 1", "number": "Number 1", "group": "Group 1", "description": "Description 1", "object_type": "asset.ProductInformation", "family": "Family 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSubscriptionRelationship(t *testing.T) {
	p := models.AssetSubscriptionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSubscriptionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSubscriptionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSubscriptionAccountRelationship(t *testing.T) {
	p := models.AssetSubscriptionAccountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSubscriptionAccountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSubscriptionAccountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSudiInfo(t *testing.T) {
	p := models.AssetSudiInfo{}
	var d = &schema.ResourceData{}
	c := `{"Signature":"Signature %d","Status":"Status %d","ClassId":"asset.SudiInfo","ObjectType":"asset.SudiInfo","Pid":"Pid %d","SerialNumber":"SerialNumber %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSudiInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSudiInfo(p, d)[0]
	expectedOp := map[string]interface{}{"signature": "Signature 1", "status": "Status 1", "class_id": "asset.SudiInfo", "object_type": "asset.SudiInfo", "pid": "Pid 1", "serial_number": "SerialNumber 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetTargetRelationship(t *testing.T) {
	p := models.AssetTargetRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetTargetRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetTargetRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosBootModeRelationship(t *testing.T) {
	p := models.BiosBootModeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosBootModeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosBootModeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosSystemBootOrderRelationship(t *testing.T) {
	p := models.BiosSystemBootOrderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosSystemBootOrderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosSystemBootOrderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosTokenSettingsRelationship(t *testing.T) {
	p := models.BiosTokenSettingsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosTokenSettingsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosTokenSettingsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosUnitRelationship(t *testing.T) {
	p := models.BiosUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosVfSelectMemoryRasConfigurationRelationship(t *testing.T) {
	p := models.BiosVfSelectMemoryRasConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosVfSelectMemoryRasConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosVfSelectMemoryRasConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBootDeviceBootModeRelationship(t *testing.T) {
	p := models.BootDeviceBootModeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBootDeviceBootModeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBootDeviceBootModeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBootDeviceBootSecurityRelationship(t *testing.T) {
	p := models.BootDeviceBootSecurityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBootDeviceBootSecurityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBootDeviceBootSecurityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchNetworkLimits(t *testing.T) {
	p := models.CapabilitySwitchNetworkLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumEthernetPortChannels":32,"MaximumFcPortChannelMembers":32,"MaximumVlans":32,"MaximumEthernetUplinkPorts":32,"MaximumPrimaryVlan":32,"ClassId":"capability.SwitchNetworkLimits","ObjectType":"capability.SwitchNetworkLimits","MaximumSecondaryVlanPerPrimary":32,"MaximumActiveTrafficMonitoringSessions":32,"MinimumActiveFans":32,"MaximumFcPortChannels":32,"MaximumPortChannelMembers":32,"MaximumSecondaryVlan":32,"MaxUncompressedPortVlanCount":32,"MaximumIgmpGroups":32,"MaxCompressedPortVlanCount":32,"MaximumVifs":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchNetworkLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchNetworkLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_ethernet_port_channels": 32, "maximum_fc_port_channel_members": 32, "maximum_vlans": 32, "maximum_ethernet_uplink_ports": 32, "maximum_primary_vlan": 32, "class_id": "capability.SwitchNetworkLimits", "object_type": "capability.SwitchNetworkLimits", "maximum_secondary_vlan_per_primary": 32, "maximum_active_traffic_monitoring_sessions": 32, "minimum_active_fans": 32, "maximum_fc_port_channels": 32, "maximum_port_channel_members": 32, "maximum_secondary_vlan": 32, "max_uncompressed_port_vlan_count": 32, "maximum_igmp_groups": 32, "max_compressed_port_vlan_count": 32, "maximum_vifs": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchStorageLimits(t *testing.T) {
	p := models.CapabilitySwitchStorageLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumZoneCount":32,"MaximumUserZoneCount":32,"MaximumVirtualFcInterfaces":32,"ClassId":"capability.SwitchStorageLimits","ObjectType":"capability.SwitchStorageLimits","MaximumVirtualFcInterfacesPerBladeServer":32,"MaximumVsans":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchStorageLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchStorageLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_zone_count": 32, "maximum_user_zone_count": 32, "maximum_virtual_fc_interfaces": 32, "class_id": "capability.SwitchStorageLimits", "object_type": "capability.SwitchStorageLimits", "maximum_virtual_fc_interfaces_per_blade_server": 32, "maximum_vsans": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchSystemLimits(t *testing.T) {
	p := models.CapabilitySwitchSystemLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumChassisCount":32,"MaximumFexPerDomain":32,"MaximumServersPerDomain":32,"ClassId":"capability.SwitchSystemLimits","ObjectType":"capability.SwitchSystemLimits"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchSystemLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchSystemLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_chassis_count": 32, "maximum_fex_per_domain": 32, "maximum_servers_per_domain": 32, "class_id": "capability.SwitchSystemLimits", "object_type": "capability.SwitchSystemLimits"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCertificatemanagementCertificateBase(t *testing.T) {
	p := models.CertificatemanagementCertificateBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"certificatemanagement.CertificateBase","Enabled":true,"IsPrivatekeySet":true,"ObjectType":"certificatemanagement.CertificateBase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCertificatemanagementCertificateBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCertificatemanagementCertificateBase(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "certificatemanagement.CertificateBase", "enabled": true, "is_privatekey_set": true, "object_type": "certificatemanagement.CertificateBase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapChassisConfigResultRelationship(t *testing.T) {
	p := models.ChassisConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapChassisConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapChassisConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapChassisProfileRelationship(t *testing.T) {
	p := models.ChassisProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapChassisProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapChassisProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudTfcOrganizationRelationship(t *testing.T) {
	p := models.CloudTfcOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudTfcOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudTfcOrganizationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommHttpProxyPolicyRelationship(t *testing.T) {
	p := models.CommHttpProxyPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommHttpProxyPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommHttpProxyPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommIpV4Interface(t *testing.T) {
	p := models.CommIpV4Interface{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"comm.IpV4Interface","ObjectType":"comm.IpV4Interface","Gateway":"Gateway %d","IpAddress":"IpAddress %d","Netmask":"Netmask %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommIpV4Interface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommIpV4Interface(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "comm.IpV4Interface", "object_type": "comm.IpV4Interface", "gateway": "Gateway 1", "ip_address": "IpAddress 1", "netmask": "Netmask 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommIpV6Interface(t *testing.T) {
	p := models.CommIpV6Interface{}
	var d = &schema.ResourceData{}
	c := `{"Prefix":"Prefix %d","ClassId":"comm.IpV6Interface","ObjectType":"comm.IpV6Interface","Gateway":"Gateway %d","IpAddress":"IpAddress %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommIpV6Interface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommIpV6Interface(p, d)[0]
	expectedOp := map[string]interface{}{"prefix": "Prefix 1", "class_id": "comm.IpV6Interface", "object_type": "comm.IpV6Interface", "gateway": "Gateway 1", "ip_address": "IpAddress 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeAlarmSummary(t *testing.T) {
	p := models.ComputeAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.AlarmSummary","ObjectType":"compute.AlarmSummary","Critical":32,"Warning":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.AlarmSummary", "object_type": "compute.AlarmSummary", "critical": 32, "warning": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeBladeRelationship(t *testing.T) {
	p := models.ComputeBladeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeBladeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeBladeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeBoardRelationship(t *testing.T) {
	p := models.ComputeBoardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeBoardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeBoardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputePersistentMemoryOperation(t *testing.T) {
	p := models.ComputePersistentMemoryOperation{}
	var d = &schema.ResourceData{}
	c := `{"AdminAction":"AdminAction %d","ClassId":"compute.PersistentMemoryOperation","ObjectType":"compute.PersistentMemoryOperation","IsSecurePassphraseSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputePersistentMemoryOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputePersistentMemoryOperation(p, d)[0]
	expectedOp := map[string]interface{}{"admin_action": "AdminAction 1", "class_id": "compute.PersistentMemoryOperation", "object_type": "compute.PersistentMemoryOperation", "is_secure_passphrase_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputePhysicalRelationship(t *testing.T) {
	p := models.ComputePhysicalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputePhysicalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputePhysicalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeRackUnitRelationship(t *testing.T) {
	p := models.ComputeRackUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeRackUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeRackUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeServerConfig(t *testing.T) {
	p := models.ComputeServerConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.ServerConfig","ObjectType":"compute.ServerConfig","AssetTag":"AssetTag %d","UserLabel":"UserLabel %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeServerConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeServerConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.ServerConfig", "object_type": "compute.ServerConfig", "asset_tag": "AssetTag 1", "user_label": "UserLabel 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStorageControllerOperation(t *testing.T) {
	p := models.ComputeStorageControllerOperation{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"compute.StorageControllerOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d","ClassId":"compute.StorageControllerOperation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStorageControllerOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStorageControllerOperation(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "compute.StorageControllerOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1", "class_id": "compute.StorageControllerOperation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStoragePhysicalDriveOperation(t *testing.T) {
	p := models.ComputeStoragePhysicalDriveOperation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.StoragePhysicalDriveOperation","ObjectType":"compute.StoragePhysicalDriveOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStoragePhysicalDriveOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStoragePhysicalDriveOperation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.StoragePhysicalDriveOperation", "object_type": "compute.StoragePhysicalDriveOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStorageVirtualDriveOperation(t *testing.T) {
	p := models.ComputeStorageVirtualDriveOperation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.StorageVirtualDriveOperation","ObjectType":"compute.StorageVirtualDriveOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStorageVirtualDriveOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStorageVirtualDriveOperation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.StorageVirtualDriveOperation", "object_type": "compute.StorageVirtualDriveOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeVmediaRelationship(t *testing.T) {
	p := models.ComputeVmediaRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeVmediaRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeVmediaRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCondAlarmSummary(t *testing.T) {
	p := models.CondAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"Critical":32,"Warning":32,"ClassId":"cond.AlarmSummary","ObjectType":"cond.AlarmSummary"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCondAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCondAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"critical": 32, "warning": 32, "class_id": "cond.AlarmSummary", "object_type": "cond.AlarmSummary"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCondHclStatusRelationship(t *testing.T) {
	p := models.CondHclStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCondHclStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCondHclStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConfigExportedItemRelationship(t *testing.T) {
	p := models.ConfigExportedItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConfigExportedItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConfigExportedItemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConfigExporterRelationship(t *testing.T) {
	p := models.ConfigExporterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConfigExporterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConfigExporterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConfigImporterRelationship(t *testing.T) {
	p := models.ConfigImporterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConfigImporterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConfigImporterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConfigMoRef(t *testing.T) {
	p := models.ConfigMoRef{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"config.MoRef","ObjectType":"config.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConfigMoRef(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConfigMoRef(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "config.MoRef", "object_type": "config.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConnectorFileChecksum(t *testing.T) {
	p := models.ConnectorFileChecksum{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"connector.FileChecksum","ObjectType":"connector.FileChecksum","HashAlgorithm":"HashAlgorithm %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConnectorFileChecksum(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConnectorFileChecksum(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "connector.FileChecksum", "object_type": "connector.FileChecksum", "hash_algorithm": "HashAlgorithm 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConnectorPlatformParamBase(t *testing.T) {
	p := models.ConnectorPlatformParamBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"connector.PlatformParamBase","ObjectType":"connector.PlatformParamBase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConnectorPlatformParamBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConnectorPlatformParamBase(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "connector.PlatformParamBase", "object_type": "connector.PlatformParamBase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentBaseRelationship(t *testing.T) {
	p := models.EquipmentBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentChassisRelationship(t *testing.T) {
	p := models.EquipmentChassisRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentChassisRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentChassisRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFanControlRelationship(t *testing.T) {
	p := models.EquipmentFanControlRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFanControlRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFanControlRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFanModuleRelationship(t *testing.T) {
	p := models.EquipmentFanModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFanModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFanModuleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFexRelationship(t *testing.T) {
	p := models.EquipmentFexRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFexRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFexRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFruRelationship(t *testing.T) {
	p := models.EquipmentFruRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFruRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFruRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentIoCardRelationship(t *testing.T) {
	p := models.EquipmentIoCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentIoCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentIoCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentIoCardBaseRelationship(t *testing.T) {
	p := models.EquipmentIoCardBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentIoCardBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentIoCardBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentLocatorLedRelationship(t *testing.T) {
	p := models.EquipmentLocatorLedRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentLocatorLedRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentLocatorLedRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentPhysicalIdentityRelationship(t *testing.T) {
	p := models.EquipmentPhysicalIdentityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentPhysicalIdentityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentPhysicalIdentityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentPsuControlRelationship(t *testing.T) {
	p := models.EquipmentPsuControlRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentPsuControlRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentPsuControlRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentRackEnclosureRelationship(t *testing.T) {
	p := models.EquipmentRackEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentRackEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentRackEnclosureRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentRackEnclosureSlotRelationship(t *testing.T) {
	p := models.EquipmentRackEnclosureSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentRackEnclosureSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentRackEnclosureSlotRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSharedIoModuleRelationship(t *testing.T) {
	p := models.EquipmentSharedIoModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSharedIoModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSharedIoModuleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSwitchCardRelationship(t *testing.T) {
	p := models.EquipmentSwitchCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSwitchCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSwitchCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSystemIoControllerRelationship(t *testing.T) {
	p := models.EquipmentSystemIoControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSystemIoControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSystemIoControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEtherPhysicalPortRelationship(t *testing.T) {
	p := models.EtherPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEtherPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEtherPhysicalPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEtherPhysicalPortBaseRelationship(t *testing.T) {
	p := models.EtherPhysicalPortBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEtherPhysicalPortBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEtherPhysicalPortBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricConfigResultRelationship(t *testing.T) {
	p := models.FabricConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkControlPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkGroupPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkGroupPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkGroupPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkGroupPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricFcNetworkPolicyRelationship(t *testing.T) {
	p := models.FabricFcNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricFcNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricFcNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricFlowControlPolicyRelationship(t *testing.T) {
	p := models.FabricFlowControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricFlowControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricFlowControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLinkAggregationPolicyRelationship(t *testing.T) {
	p := models.FabricLinkAggregationPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLinkAggregationPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLinkAggregationPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLinkControlPolicyRelationship(t *testing.T) {
	p := models.FabricLinkControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLinkControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLinkControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLldpSettings(t *testing.T) {
	p := models.FabricLldpSettings{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"fabric.LldpSettings","ReceiveEnabled":true,"TransmitEnabled":true,"ClassId":"fabric.LldpSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLldpSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLldpSettings(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "fabric.LldpSettings", "receive_enabled": true, "transmit_enabled": true, "class_id": "fabric.LldpSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricMacAgingSettings(t *testing.T) {
	p := models.FabricMacAgingSettings{}
	var d = &schema.ResourceData{}
	c := `{"MacAgingTime":32,"ClassId":"fabric.MacAgingSettings","ObjectType":"fabric.MacAgingSettings","MacAgingOption":"MacAgingOption %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricMacAgingSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricMacAgingSettings(p, d)[0]
	expectedOp := map[string]interface{}{"mac_aging_time": 32, "class_id": "fabric.MacAgingSettings", "object_type": "fabric.MacAgingSettings", "mac_aging_option": "MacAgingOption 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricMulticastPolicyRelationship(t *testing.T) {
	p := models.FabricMulticastPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricMulticastPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricMulticastPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricPortPolicyRelationship(t *testing.T) {
	p := models.FabricPortPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricPortPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricPortPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricSwitchClusterProfileRelationship(t *testing.T) {
	p := models.FabricSwitchClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricSwitchClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricSwitchClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricSwitchProfileRelationship(t *testing.T) {
	p := models.FabricSwitchProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricSwitchProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricSwitchProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricUdldGlobalSettings(t *testing.T) {
	p := models.FabricUdldGlobalSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.UdldGlobalSettings","ObjectType":"fabric.UdldGlobalSettings","MessageInterval":32,"RecoveryAction":"RecoveryAction %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricUdldGlobalSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricUdldGlobalSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.UdldGlobalSettings", "object_type": "fabric.UdldGlobalSettings", "message_interval": 32, "recovery_action": "RecoveryAction 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricUdldSettings(t *testing.T) {
	p := models.FabricUdldSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.UdldSettings","ObjectType":"fabric.UdldSettings","AdminState":"AdminState %d","Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricUdldSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricUdldSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.UdldSettings", "object_type": "fabric.UdldSettings", "admin_state": "AdminState 1", "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricVlanSettings(t *testing.T) {
	p := models.FabricVlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.VlanSettings","ObjectType":"fabric.VlanSettings","AllowedVlans":"AllowedVlans %d","NativeVlan":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricVlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricVlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.VlanSettings", "object_type": "fabric.VlanSettings", "allowed_vlans": "AllowedVlans 1", "native_vlan": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolBlock(t *testing.T) {
	p := models.FcpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ClassId":"fcpool.Block","ObjectType":"fcpool.Block","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolBlock(p, d)[0]
	expectedOp := map[string]interface{}{"from": "From 1", "to": "To 1", "class_id": "fcpool.Block", "object_type": "fcpool.Block", "size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolFcBlockRelationship(t *testing.T) {
	p := models.FcpoolFcBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolFcBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolFcBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolLeaseRelationship(t *testing.T) {
	p := models.FcpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolPoolRelationship(t *testing.T) {
	p := models.FcpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolPoolMemberRelationship(t *testing.T) {
	p := models.FcpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolUniverseRelationship(t *testing.T) {
	p := models.FcpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareBaseDistributableRelationship(t *testing.T) {
	p := models.FirmwareBaseDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareBaseDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareBaseDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareDirectDownload(t *testing.T) {
	p := models.FirmwareDirectDownload{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"firmware.DirectDownload","ClassId":"firmware.DirectDownload","Username":"Username %d","IsPasswordSet":true,"Upgradeoption":"Upgradeoption %d","ImageSource":"ImageSource %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareDirectDownload(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareDirectDownload(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "firmware.DirectDownload", "class_id": "firmware.DirectDownload", "username": "Username 1", "is_password_set": true, "upgradeoption": "Upgradeoption 1", "image_source": "ImageSource 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareDistributableRelationship(t *testing.T) {
	p := models.FirmwareDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareNetworkShare(t *testing.T) {
	p := models.FirmwareNetworkShare{}
	var d = &schema.ResourceData{}
	c := `{"Upgradeoption":"Upgradeoption %d","ClassId":"firmware.NetworkShare","MapType":"MapType %d","Username":"Username %d","ObjectType":"firmware.NetworkShare","IsPasswordSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareNetworkShare(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareNetworkShare(p, d)[0]
	expectedOp := map[string]interface{}{"upgradeoption": "Upgradeoption 1", "class_id": "firmware.NetworkShare", "map_type": "MapType 1", "username": "Username 1", "object_type": "firmware.NetworkShare", "is_password_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareRunningFirmwareRelationship(t *testing.T) {
	p := models.FirmwareRunningFirmwareRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareRunningFirmwareRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareRunningFirmwareRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareServerConfigurationUtilityDistributableRelationship(t *testing.T) {
	p := models.FirmwareServerConfigurationUtilityDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeBaseRelationship(t *testing.T) {
	p := models.FirmwareUpgradeBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeImpactStatusRelationship(t *testing.T) {
	p := models.FirmwareUpgradeImpactStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeImpactStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeImpactStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeStatusRelationship(t *testing.T) {
	p := models.FirmwareUpgradeStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastCatalogRelationship(t *testing.T) {
	p := models.ForecastCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastDefinitionRelationship(t *testing.T) {
	p := models.ForecastDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastInstanceRelationship(t *testing.T) {
	p := models.ForecastInstanceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastInstanceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastInstanceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastModel(t *testing.T) {
	p := models.ForecastModel{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"forecast.Model","Accuracy":32.000000,"ModelType":"ModelType %d","ClassId":"forecast.Model"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastModel(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastModel(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "forecast.Model", "accuracy": 32.000000, "model_type": "ModelType 1", "class_id": "forecast.Model"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapGraphicsCardRelationship(t *testing.T) {
	p := models.GraphicsCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapGraphicsCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapGraphicsCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHclOperatingSystemRelationship(t *testing.T) {
	p := models.HclOperatingSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHclOperatingSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHclOperatingSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHclOperatingSystemVendorRelationship(t *testing.T) {
	p := models.HclOperatingSystemVendorRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHclOperatingSystemVendorRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHclOperatingSystemVendorRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAlarmSummary(t *testing.T) {
	p := models.HyperflexAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.AlarmSummary","ObjectType":"hyperflex.AlarmSummary","Warning":32,"Critical":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.AlarmSummary", "object_type": "hyperflex.AlarmSummary", "warning": 32, "critical": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAppCatalogRelationship(t *testing.T) {
	p := models.HyperflexAppCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAppCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAppCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAppSettingConstraint(t *testing.T) {
	p := models.HyperflexAppSettingConstraint{}
	var d = &schema.ResourceData{}
	c := `{"HypervisorType":"HypervisorType %d","MgmtPlatform":"MgmtPlatform %d","ServerModel":"ServerModel %d","ClassId":"hyperflex.AppSettingConstraint","ObjectType":"hyperflex.AppSettingConstraint","HxdpVersion":"HxdpVersion %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAppSettingConstraint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAppSettingConstraint(p, d)[0]
	expectedOp := map[string]interface{}{"hypervisor_type": "HypervisorType 1", "mgmt_platform": "MgmtPlatform 1", "server_model": "ServerModel 1", "class_id": "hyperflex.AppSettingConstraint", "object_type": "hyperflex.AppSettingConstraint", "hxdp_version": "HxdpVersion 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAutoSupportPolicyRelationship(t *testing.T) {
	p := models.HyperflexAutoSupportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAutoSupportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAutoSupportPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexBackupClusterRelationship(t *testing.T) {
	p := models.HyperflexBackupClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexBackupClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexBackupClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexBaseClusterRelationship(t *testing.T) {
	p := models.HyperflexBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexBaseClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexBondState(t *testing.T) {
	p := models.HyperflexBondState{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.BondState","ObjectType":"hyperflex.BondState","Mode":"Mode %d","ActiveSlave":"ActiveSlave %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexBondState(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexBondState(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.BondState", "object_type": "hyperflex.BondState", "mode": "Mode 1", "active_slave": "ActiveSlave 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexCiscoHypervisorManagerRelationship(t *testing.T) {
	p := models.HyperflexCiscoHypervisorManagerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexCiscoHypervisorManagerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexCiscoHypervisorManagerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterRelationship(t *testing.T) {
	p := models.HyperflexClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterNetworkPolicyRelationship(t *testing.T) {
	p := models.HyperflexClusterNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterProfileRelationship(t *testing.T) {
	p := models.HyperflexClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexClusterStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexConfigResultRelationship(t *testing.T) {
	p := models.HyperflexConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexDataProtectionPeerRelationship(t *testing.T) {
	p := models.HyperflexDataProtectionPeerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexDataProtectionPeerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexDataProtectionPeerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexDiskStatus(t *testing.T) {
	p := models.HyperflexDiskStatus{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.DiskStatus","DownloadPercentage":"DownloadPercentage %d","State":"State %d","VolumeHandle":"VolumeHandle %d","VolumeName":"VolumeName %d","ClassId":"hyperflex.DiskStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexDiskStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexDiskStatus(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "hyperflex.DiskStatus", "download_percentage": "DownloadPercentage 1", "state": "State 1", "volume_handle": "VolumeHandle 1", "volume_name": "VolumeName 1", "class_id": "hyperflex.DiskStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexEntityReference(t *testing.T) {
	p := models.HyperflexEntityReference{}
	var d = &schema.ResourceData{}
	c := `{"Confignum":32,"Id":"Id %d","Idtype":"Idtype %d","Name":"Name %d","Type":"Type %d","ClassId":"hyperflex.EntityReference","ObjectType":"hyperflex.EntityReference"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexEntityReference(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexEntityReference(p, d)[0]
	expectedOp := map[string]interface{}{"confignum": 32, "id": "Id 1", "idtype": "Idtype 1", "name": "Name 1", "type": "Type 1", "class_id": "hyperflex.EntityReference", "object_type": "hyperflex.EntityReference"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexErrorStack(t *testing.T) {
	p := models.HyperflexErrorStack{}
	var d = &schema.ResourceData{}
	c := `{"MessageId":32,"Message":"Message %d","ObjectType":"hyperflex.ErrorStack","ClassId":"hyperflex.ErrorStack"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexErrorStack(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexErrorStack(p, d)[0]
	expectedOp := map[string]interface{}{"message_id": 32, "message": "Message 1", "object_type": "hyperflex.ErrorStack", "class_id": "hyperflex.ErrorStack"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexExtFcStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexExtFcStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexExtFcStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexExtFcStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexExtIscsiStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexExtIscsiStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexExtIscsiStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexExtIscsiStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexFeatureLimitExternalRelationship(t *testing.T) {
	p := models.HyperflexFeatureLimitExternalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexFeatureLimitExternalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexFeatureLimitExternalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexFeatureLimitInternalRelationship(t *testing.T) {
	p := models.HyperflexFeatureLimitInternalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexFeatureLimitInternalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexFeatureLimitInternalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthRelationship(t *testing.T) {
	p := models.HyperflexHealthRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthCheckDefinitionRelationship(t *testing.T) {
	p := models.HyperflexHealthCheckDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthCheckDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthCheckDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthCheckScriptInfo(t *testing.T) {
	p := models.HyperflexHealthCheckScriptInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HealthCheckScriptInfo","ScriptInput":"ScriptInput %d","AggregateScriptName":"AggregateScriptName %d","HyperflexVersion":"HyperflexVersion %d","ScriptExecuteLocation":"ScriptExecuteLocation %d","ScriptName":"ScriptName %d","Version":"Version %d","ObjectType":"hyperflex.HealthCheckScriptInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthCheckScriptInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthCheckScriptInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HealthCheckScriptInfo", "script_input": "ScriptInput 1", "aggregate_script_name": "AggregateScriptName 1", "hyperflex_version": "HyperflexVersion 1", "script_execute_location": "ScriptExecuteLocation 1", "script_name": "ScriptName 1", "nr_version": "Version 1", "object_type": "hyperflex.HealthCheckScriptInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxLicenseAuthorizationDetailsDt(t *testing.T) {
	p := models.HyperflexHxLicenseAuthorizationDetailsDt{}
	var d = &schema.ResourceData{}
	c := `{"EvaluationPeriodExpiredAt":"EvaluationPeriodExpiredAt %d","CommunicationDeadlineDate":"CommunicationDeadlineDate %d","ClassId":"hyperflex.HxLicenseAuthorizationDetailsDt","ObjectType":"hyperflex.HxLicenseAuthorizationDetailsDt","LastCommunicationAttemptDate":"LastCommunicationAttemptDate %d","NextCommunicationAttemptDate":"NextCommunicationAttemptDate %d","Status":"Status %d","EvaluationPeriodRemaining":"EvaluationPeriodRemaining %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxLicenseAuthorizationDetailsDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxLicenseAuthorizationDetailsDt(p, d)[0]
	expectedOp := map[string]interface{}{"evaluation_period_expired_at": "EvaluationPeriodExpiredAt 1", "communication_deadline_date": "CommunicationDeadlineDate 1", "class_id": "hyperflex.HxLicenseAuthorizationDetailsDt", "object_type": "hyperflex.HxLicenseAuthorizationDetailsDt", "last_communication_attempt_date": "LastCommunicationAttemptDate 1", "next_communication_attempt_date": "NextCommunicationAttemptDate 1", "status": "Status 1", "evaluation_period_remaining": "EvaluationPeriodRemaining 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxNetworkAddressDt(t *testing.T) {
	p := models.HyperflexHxNetworkAddressDt{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.HxNetworkAddressDt","ClassId":"hyperflex.HxNetworkAddressDt","Address":"Address %d","Fqdn":"Fqdn %d","Ip":"Ip %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxNetworkAddressDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxNetworkAddressDt(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "hyperflex.HxNetworkAddressDt", "class_id": "hyperflex.HxNetworkAddressDt", "address": "Address 1", "fqdn": "Fqdn 1", "ip": "Ip 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxPlatformDatastoreConfigDt(t *testing.T) {
	p := models.HyperflexHxPlatformDatastoreConfigDt{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","NumMirrors":32,"ClassId":"hyperflex.HxPlatformDatastoreConfigDt","ProvisionedCapacity":32,"SystemDatastore":true,"DataBlockSize":32,"ObjectType":"hyperflex.HxPlatformDatastoreConfigDt","NumStripesForLargeFiles":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxPlatformDatastoreConfigDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxPlatformDatastoreConfigDt(p, d)[0]
	expectedOp := map[string]interface{}{"name": "Name 1", "num_mirrors": 32, "class_id": "hyperflex.HxPlatformDatastoreConfigDt", "provisioned_capacity": 32, "system_datastore": true, "data_block_size": 32, "object_type": "hyperflex.HxPlatformDatastoreConfigDt", "num_stripes_for_large_files": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxRegistrationDetailsDt(t *testing.T) {
	p := models.HyperflexHxRegistrationDetailsDt{}
	var d = &schema.ResourceData{}
	c := `{"InitialRegistrationDate":"InitialRegistrationDate %d","NextRenewalAttemptDate":"NextRenewalAttemptDate %d","RegistrationExpireDate":"RegistrationExpireDate %d","VirtualAccount":"VirtualAccount %d","LastRenewalAttemptDate":"LastRenewalAttemptDate %d","RegistrationFailedReason":"RegistrationFailedReason %d","Status":"Status %d","OutOfComplianceStartDate":"OutOfComplianceStartDate %d","SmartAccount":"SmartAccount %d","ClassId":"hyperflex.HxRegistrationDetailsDt","ObjectType":"hyperflex.HxRegistrationDetailsDt"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxRegistrationDetailsDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxRegistrationDetailsDt(p, d)[0]
	expectedOp := map[string]interface{}{"initial_registration_date": "InitialRegistrationDate 1", "next_renewal_attempt_date": "NextRenewalAttemptDate 1", "registration_expire_date": "RegistrationExpireDate 1", "virtual_account": "VirtualAccount 1", "last_renewal_attempt_date": "LastRenewalAttemptDate 1", "registration_failed_reason": "RegistrationFailedReason 1", "status": "Status 1", "out_of_compliance_start_date": "OutOfComplianceStartDate 1", "smart_account": "SmartAccount 1", "class_id": "hyperflex.HxRegistrationDetailsDt", "object_type": "hyperflex.HxRegistrationDetailsDt"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxResiliencyInfoDt(t *testing.T) {
	p := models.HyperflexHxResiliencyInfoDt{}
	var d = &schema.ResourceData{}
	c := `{"ResiliencyState":"ResiliencyState %d","SsdFailuresTolerable":32,"HddFailuresTolerable":32,"PolicyCompliance":"PolicyCompliance %d","ClassId":"hyperflex.HxResiliencyInfoDt","ObjectType":"hyperflex.HxResiliencyInfoDt","DataReplicationFactor":"DataReplicationFactor %d","NodeFailuresTolerable":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxResiliencyInfoDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxResiliencyInfoDt(p, d)[0]
	expectedOp := map[string]interface{}{"resiliency_state": "ResiliencyState 1", "ssd_failures_tolerable": 32, "hdd_failures_tolerable": 32, "policy_compliance": "PolicyCompliance 1", "class_id": "hyperflex.HxResiliencyInfoDt", "object_type": "hyperflex.HxResiliencyInfoDt", "data_replication_factor": "DataReplicationFactor 1", "node_failures_tolerable": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxSiteDt(t *testing.T) {
	p := models.HyperflexHxSiteDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxSiteDt","ObjectType":"hyperflex.HxSiteDt","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxSiteDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxSiteDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxSiteDt", "object_type": "hyperflex.HxSiteDt", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxUuIdDt(t *testing.T) {
	p := models.HyperflexHxUuIdDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxUuIdDt","ObjectType":"hyperflex.HxUuIdDt","Uuid":"Uuid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxUuIdDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxUuIdDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxUuIdDt", "object_type": "hyperflex.HxUuIdDt", "uuid": "Uuid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapClusterRelationship(t *testing.T) {
	p := models.HyperflexHxapClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapDvUplinkRelationship(t *testing.T) {
	p := models.HyperflexHxapDvUplinkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapDvUplinkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapDvUplinkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapDvswitchRelationship(t *testing.T) {
	p := models.HyperflexHxapDvswitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapDvswitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapDvswitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapHostRelationship(t *testing.T) {
	p := models.HyperflexHxapHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapNetworkRelationship(t *testing.T) {
	p := models.HyperflexHxapNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxapVirtualMachineRelationship(t *testing.T) {
	p := models.HyperflexHxapVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxapVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxapVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexIpAddrRange(t *testing.T) {
	p := models.HyperflexIpAddrRange{}
	var d = &schema.ResourceData{}
	c := `{"StartAddr":"StartAddr %d","ClassId":"hyperflex.IpAddrRange","ObjectType":"hyperflex.IpAddrRange","EndAddr":"EndAddr %d","Gateway":"Gateway %d","Netmask":"Netmask %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexIpAddrRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexIpAddrRange(p, d)[0]
	expectedOp := map[string]interface{}{"start_addr": "StartAddr 1", "class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange", "end_addr": "EndAddr 1", "gateway": "Gateway 1", "netmask": "Netmask 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLicenseRelationship(t *testing.T) {
	p := models.HyperflexLicenseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLicenseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLicenseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLocalCredentialPolicyRelationship(t *testing.T) {
	p := models.HyperflexLocalCredentialPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLocalCredentialPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLocalCredentialPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLogicalAvailabilityZone(t *testing.T) {
	p := models.HyperflexLogicalAvailabilityZone{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.LogicalAvailabilityZone","ObjectType":"hyperflex.LogicalAvailabilityZone","AutoConfig":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLogicalAvailabilityZone(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLogicalAvailabilityZone(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.LogicalAvailabilityZone", "object_type": "hyperflex.LogicalAvailabilityZone", "auto_config": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexMacAddrPrefixRange(t *testing.T) {
	p := models.HyperflexMacAddrPrefixRange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.MacAddrPrefixRange","ObjectType":"hyperflex.MacAddrPrefixRange","EndAddr":"EndAddr %d","StartAddr":"StartAddr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexMacAddrPrefixRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexMacAddrPrefixRange(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.MacAddrPrefixRange", "object_type": "hyperflex.MacAddrPrefixRange", "end_addr": "EndAddr 1", "start_addr": "StartAddr 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNamedVlan(t *testing.T) {
	p := models.HyperflexNamedVlan{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.NamedVlan","ObjectType":"hyperflex.NamedVlan","Name":"Name %d","VlanId":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNamedVlan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNamedVlan(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "name": "Name 1", "vlan_id": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNamedVsan(t *testing.T) {
	p := models.HyperflexNamedVsan{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","VsanId":32,"ClassId":"hyperflex.NamedVsan","ObjectType":"hyperflex.NamedVsan"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNamedVsan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNamedVsan(p, d)[0]
	expectedOp := map[string]interface{}{"name": "Name 1", "vsan_id": 32, "class_id": "hyperflex.NamedVsan", "object_type": "hyperflex.NamedVsan"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNodeConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexNodeConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNodeConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNodeConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexProxySettingPolicyRelationship(t *testing.T) {
	p := models.HyperflexProxySettingPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexProxySettingPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexProxySettingPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexReplicationPeerInfo(t *testing.T) {
	p := models.HyperflexReplicationPeerInfo{}
	var d = &schema.ResourceData{}
	c := `{"Mcip":"Mcip %d","StatusDetails":"StatusDetails %d","ObjectType":"hyperflex.ReplicationPeerInfo","Dcip":"Dcip %d","ReplCip":"ReplCip %d","ClassId":"hyperflex.ReplicationPeerInfo","Status":"Status %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexReplicationPeerInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexReplicationPeerInfo(p, d)[0]
	expectedOp := map[string]interface{}{"mcip": "Mcip 1", "status_details": "StatusDetails 1", "object_type": "hyperflex.ReplicationPeerInfo", "dcip": "Dcip 1", "repl_cip": "ReplCip 1", "class_id": "hyperflex.ReplicationPeerInfo", "status": "Status 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexReplicationSchedule(t *testing.T) {
	p := models.HyperflexReplicationSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.ReplicationSchedule","BackupInterval":32,"ClassId":"hyperflex.ReplicationSchedule"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexReplicationSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexReplicationSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "hyperflex.ReplicationSchedule", "backup_interval": 32, "class_id": "hyperflex.ReplicationSchedule"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexServerFirmwareVersionRelationship(t *testing.T) {
	p := models.HyperflexServerFirmwareVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexServerFirmwareVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexServerFirmwareVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexServerModelRelationship(t *testing.T) {
	p := models.HyperflexServerModelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexServerModelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexServerModelRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareDistributionEntryRelationship(t *testing.T) {
	p := models.HyperflexSoftwareDistributionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareDistributionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareDistributionEntryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareDistributionVersionRelationship(t *testing.T) {
	p := models.HyperflexSoftwareDistributionVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareDistributionVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareDistributionVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareVersionPolicyRelationship(t *testing.T) {
	p := models.HyperflexSoftwareVersionPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareVersionPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareVersionPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexStorageContainerRelationship(t *testing.T) {
	p := models.HyperflexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexStorageContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSummary(t *testing.T) {
	p := models.HyperflexSummary{}
	var d = &schema.ResourceData{}
	c := `{"DeduplicationSavings":32.000000,"ClassId":"hyperflex.Summary","Name":"Name %d","State":"State %d","Downtime":"Downtime %d","ActiveNodes":"ActiveNodes %d","ResiliencyDetailsSize":32,"DataReplicationCompliance":"DataReplicationCompliance %d","FreeCapacity":32,"Address":"Address %d","UsedCapacity":32,"CompressionSavings":32.000000,"SpaceStatus":"SpaceStatus %d","TotalCapacity":32,"Uptime":"Uptime %d","DataReplicationFactor":"DataReplicationFactor %d","ObjectType":"hyperflex.Summary","ClusterAccessPolicy":"ClusterAccessPolicy %d","Boottime":32,"TotalSavings":32.000000}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSummary(p, d)[0]
	expectedOp := map[string]interface{}{"deduplication_savings": 32.000000, "class_id": "hyperflex.Summary", "name": "Name 1", "state": "State 1", "downtime": "Downtime 1", "active_nodes": "ActiveNodes 1", "resiliency_details_size": 32, "data_replication_compliance": "DataReplicationCompliance 1", "free_capacity": 32, "address": "Address 1", "used_capacity": 32, "compression_savings": 32.000000, "space_status": "SpaceStatus 1", "total_capacity": 32, "uptime": "Uptime 1", "data_replication_factor": "DataReplicationFactor 1", "object_type": "hyperflex.Summary", "cluster_access_policy": "ClusterAccessPolicy 1", "boottime": 32, "total_savings": 32.000000}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSysConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexSysConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSysConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSysConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexUcsmConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexUcsmConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexUcsmConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexUcsmConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVcenterConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexVcenterConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVcenterConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVcenterConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVirtualMachine(t *testing.T) {
	p := models.HyperflexVirtualMachine{}
	var d = &schema.ResourceData{}
	c := `{"StatusCode":"StatusCode %d","Uuid":"Uuid %d","ObjectType":"hyperflex.VirtualMachine","ClassId":"hyperflex.VirtualMachine"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVirtualMachine(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVirtualMachine(p, d)[0]
	expectedOp := map[string]interface{}{"status_code": "StatusCode 1", "uuid": "Uuid 1", "object_type": "hyperflex.VirtualMachine", "class_id": "hyperflex.VirtualMachine"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVmBackupInfoRelationship(t *testing.T) {
	p := models.HyperflexVmBackupInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVmBackupInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVmBackupInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVmSnapshotInfoRelationship(t *testing.T) {
	p := models.HyperflexVmSnapshotInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVmSnapshotInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVmSnapshotInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexWwxnPrefixRange(t *testing.T) {
	p := models.HyperflexWwxnPrefixRange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.WwxnPrefixRange","ObjectType":"hyperflex.WwxnPrefixRange","StartAddr":"StartAddr %d","EndAddr":"EndAddr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexWwxnPrefixRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexWwxnPrefixRange(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.WwxnPrefixRange", "object_type": "hyperflex.WwxnPrefixRange", "start_addr": "StartAddr 1", "end_addr": "EndAddr 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasLicenseInfoRelationship(t *testing.T) {
	p := models.IaasLicenseInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasLicenseInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasLicenseInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasUcsdInfoRelationship(t *testing.T) {
	p := models.IaasUcsdInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasUcsdInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasUcsdInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasUcsdManagedInfraRelationship(t *testing.T) {
	p := models.IaasUcsdManagedInfraRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasUcsdManagedInfraRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasUcsdManagedInfraRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamAccountRelationship(t *testing.T) {
	p := models.IamAccountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamAccountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamAccountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamAppRegistrationRelationship(t *testing.T) {
	p := models.IamAppRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamAppRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamAppRegistrationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamCertificateRelationship(t *testing.T) {
	p := models.IamCertificateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamCertificateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamCertificateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamCertificateRequestRelationship(t *testing.T) {
	p := models.IamCertificateRequestRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamCertificateRequestRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamCertificateRequestRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamClientMeta(t *testing.T) {
	p := models.IamClientMeta{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"iam.ClientMeta","DeviceModel":"DeviceModel %d","UserAgent":"UserAgent %d","ClassId":"iam.ClientMeta"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamClientMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamClientMeta(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "iam.ClientMeta", "device_model": "DeviceModel 1", "user_agent": "UserAgent 1", "class_id": "iam.ClientMeta"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamDomainGroupRelationship(t *testing.T) {
	p := models.IamDomainGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamDomainGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamDomainGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointPasswordProperties(t *testing.T) {
	p := models.IamEndPointPasswordProperties{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"iam.EndPointPasswordProperties","ClassId":"iam.EndPointPasswordProperties","PasswordHistory":32,"EnablePasswordExpiry":true,"ForceSendPassword":true,"EnforceStrongPassword":true,"PasswordExpiryDuration":32,"GracePeriod":32,"NotificationPeriod":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointPasswordProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointPasswordProperties(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "iam.EndPointPasswordProperties", "class_id": "iam.EndPointPasswordProperties", "password_history": 32, "enable_password_expiry": true, "force_send_password": true, "enforce_strong_password": true, "password_expiry_duration": 32, "grace_period": 32, "notification_period": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointUserRelationship(t *testing.T) {
	p := models.IamEndPointUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointUserRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointUserPolicyRelationship(t *testing.T) {
	p := models.IamEndPointUserPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointUserPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointUserPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIdpRelationship(t *testing.T) {
	p := models.IamIdpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIdpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIdpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIdpReferenceRelationship(t *testing.T) {
	p := models.IamIdpReferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIdpReferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIdpReferenceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIpAccessManagementRelationship(t *testing.T) {
	p := models.IamIpAccessManagementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIpAccessManagementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIpAccessManagementRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapBaseProperties(t *testing.T) {
	p := models.IamLdapBaseProperties{}
	var d = &schema.ResourceData{}
	c := `{"EnableEncryption":true,"BindMethod":"BindMethod %d","BaseDn":"BaseDn %d","IsPasswordSet":true,"ObjectType":"iam.LdapBaseProperties","NestedGroupSearchDepth":32,"Domain":"Domain %d","Filter":"Filter %d","GroupAttribute":"GroupAttribute %d","BindDn":"BindDn %d","Timeout":32,"ClassId":"iam.LdapBaseProperties","Attribute":"Attribute %d","EnableGroupAuthorization":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapBaseProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapBaseProperties(p, d)[0]
	expectedOp := map[string]interface{}{"enable_encryption": true, "bind_method": "BindMethod 1", "base_dn": "BaseDn 1", "is_password_set": true, "object_type": "iam.LdapBaseProperties", "nested_group_search_depth": 32, "domain": "Domain 1", "filter": "Filter 1", "group_attribute": "GroupAttribute 1", "bind_dn": "BindDn 1", "timeout": 32, "class_id": "iam.LdapBaseProperties", "attribute": "Attribute 1", "enable_group_authorization": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapDnsParameters(t *testing.T) {
	p := models.IamLdapDnsParameters{}
	var d = &schema.ResourceData{}
	c := `{"SearchDomain":"SearchDomain %d","SearchForest":"SearchForest %d","Source":"Source %d","ClassId":"iam.LdapDnsParameters","ObjectType":"iam.LdapDnsParameters"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapDnsParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapDnsParameters(p, d)[0]
	expectedOp := map[string]interface{}{"search_domain": "SearchDomain 1", "search_forest": "SearchForest 1", "nr_source": "Source 1", "class_id": "iam.LdapDnsParameters", "object_type": "iam.LdapDnsParameters"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapPolicyRelationship(t *testing.T) {
	p := models.IamLdapPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLocalUserPasswordRelationship(t *testing.T) {
	p := models.IamLocalUserPasswordRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLocalUserPasswordRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLocalUserPasswordRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamPermissionRelationship(t *testing.T) {
	p := models.IamPermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamPermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamPermissionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamPrivateKeySpecRelationship(t *testing.T) {
	p := models.IamPrivateKeySpecRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamPrivateKeySpecRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamPrivateKeySpecRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamQualifierRelationship(t *testing.T) {
	p := models.IamQualifierRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamQualifierRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamQualifierRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamResourceLimitsRelationship(t *testing.T) {
	p := models.IamResourceLimitsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamResourceLimitsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamResourceLimitsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSecurityHolderRelationship(t *testing.T) {
	p := models.IamSecurityHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSecurityHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSecurityHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamServiceProviderRelationship(t *testing.T) {
	p := models.IamServiceProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamServiceProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamServiceProviderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSessionRelationship(t *testing.T) {
	p := models.IamSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSessionLimitsRelationship(t *testing.T) {
	p := models.IamSessionLimitsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSessionLimitsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSessionLimitsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSystemRelationship(t *testing.T) {
	p := models.IamSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamUserRelationship(t *testing.T) {
	p := models.IamUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamUserRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamUserGroupRelationship(t *testing.T) {
	p := models.IamUserGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamUserGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamUserGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInfraHardwareInfo(t *testing.T) {
	p := models.InfraHardwareInfo{}
	var d = &schema.ResourceData{}
	c := `{"CpuCores":32,"CpuSpeed":32,"MemorySize":32,"ObjectType":"infra.HardwareInfo","ClassId":"infra.HardwareInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInfraHardwareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInfraHardwareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"cpu_cores": 32, "cpu_speed": 32, "memory_size": 32, "object_type": "infra.HardwareInfo", "class_id": "infra.HardwareInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryBaseRelationship(t *testing.T) {
	p := models.InventoryBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryDeviceInfoRelationship(t *testing.T) {
	p := models.InventoryDeviceInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryDeviceInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryDeviceInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryGenericInventoryHolderRelationship(t *testing.T) {
	p := models.InventoryGenericInventoryHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryGenericInventoryHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryGenericInventoryHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolBlockLeaseRelationship(t *testing.T) {
	p := models.IppoolBlockLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolBlockLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolBlockLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpLeaseRelationship(t *testing.T) {
	p := models.IppoolIpLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV4Block(t *testing.T) {
	p := models.IppoolIpV4Block{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ippool.IpV4Block","Size":32,"From":"From %d","To":"To %d","ClassId":"ippool.IpV4Block"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV4Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV4Block(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "ippool.IpV4Block", "size": 32, "from": "From 1", "to": "To 1", "class_id": "ippool.IpV4Block"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV4Config(t *testing.T) {
	p := models.IppoolIpV4Config{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"ippool.IpV4Config","ObjectType":"ippool.IpV4Config","PrimaryDns":"PrimaryDns %d","SecondaryDns":"SecondaryDns %d","Gateway":"Gateway %d","Netmask":"Netmask %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV4Config(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV4Config(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "ippool.IpV4Config", "object_type": "ippool.IpV4Config", "primary_dns": "PrimaryDns 1", "secondary_dns": "SecondaryDns 1", "gateway": "Gateway 1", "netmask": "Netmask 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV6Block(t *testing.T) {
	p := models.IppoolIpV6Block{}
	var d = &schema.ResourceData{}
	c := `{"Size":32,"To":"To %d","From":"From %d","ClassId":"ippool.IpV6Block","ObjectType":"ippool.IpV6Block"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV6Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV6Block(p, d)[0]
	expectedOp := map[string]interface{}{"size": 32, "to": "To 1", "from": "From 1", "class_id": "ippool.IpV6Block", "object_type": "ippool.IpV6Block"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV6Config(t *testing.T) {
	p := models.IppoolIpV6Config{}
	var d = &schema.ResourceData{}
	c := `{"PrimaryDns":"PrimaryDns %d","SecondaryDns":"SecondaryDns %d","Gateway":"Gateway %d","Prefix":32,"ClassId":"ippool.IpV6Config","ObjectType":"ippool.IpV6Config"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV6Config(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV6Config(p, d)[0]
	expectedOp := map[string]interface{}{"primary_dns": "PrimaryDns 1", "secondary_dns": "SecondaryDns 1", "gateway": "Gateway 1", "prefix": 32, "class_id": "ippool.IpV6Config", "object_type": "ippool.IpV6Config"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolPoolRelationship(t *testing.T) {
	p := models.IppoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolPoolMemberRelationship(t *testing.T) {
	p := models.IppoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolShadowBlockRelationship(t *testing.T) {
	p := models.IppoolShadowBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolShadowBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolShadowBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolShadowPoolRelationship(t *testing.T) {
	p := models.IppoolShadowPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolShadowPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolShadowPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolUniverseRelationship(t *testing.T) {
	p := models.IppoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolBlockRelationship(t *testing.T) {
	p := models.IqnpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolIqnSuffixBlock(t *testing.T) {
	p := models.IqnpoolIqnSuffixBlock{}
	var d = &schema.ResourceData{}
	c := `{"To":32,"From":32,"Suffix":"Suffix %d","ClassId":"iqnpool.IqnSuffixBlock","ObjectType":"iqnpool.IqnSuffixBlock","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolIqnSuffixBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolIqnSuffixBlock(p, d)[0]
	expectedOp := map[string]interface{}{"to": 32, "from": 32, "suffix": "Suffix 1", "class_id": "iqnpool.IqnSuffixBlock", "object_type": "iqnpool.IqnSuffixBlock", "size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolLeaseRelationship(t *testing.T) {
	p := models.IqnpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolPoolRelationship(t *testing.T) {
	p := models.IqnpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolPoolMemberRelationship(t *testing.T) {
	p := models.IqnpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolUniverseRelationship(t *testing.T) {
	p := models.IqnpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAciCniProfileRelationship(t *testing.T) {
	p := models.KubernetesAciCniProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAciCniProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAciCniProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesActionInfo(t *testing.T) {
	p := models.KubernetesActionInfo{}
	var d = &schema.ResourceData{}
	c := `{"FailureReason":"FailureReason %d","Name":"Name %d","Status":"Status %d","ClassId":"kubernetes.ActionInfo","ObjectType":"kubernetes.ActionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesActionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesActionInfo(p, d)[0]
	expectedOp := map[string]interface{}{"failure_reason": "FailureReason 1", "name": "Name 1", "status": "Status 1", "class_id": "kubernetes.ActionInfo", "object_type": "kubernetes.ActionInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAddonConfiguration(t *testing.T) {
	p := models.KubernetesAddonConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ReleaseNamespace":"ReleaseNamespace %d","InstallStrategy":"InstallStrategy %d","Overrides":"Overrides %d","ClassId":"kubernetes.AddonConfiguration","ObjectType":"kubernetes.AddonConfiguration","UpgradeStrategy":"UpgradeStrategy %d","ReleaseName":"ReleaseName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAddonConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAddonConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"release_namespace": "ReleaseNamespace 1", "install_strategy": "InstallStrategy 1", "overrides": "Overrides 1", "class_id": "kubernetes.AddonConfiguration", "object_type": "kubernetes.AddonConfiguration", "upgrade_strategy": "UpgradeStrategy 1", "release_name": "ReleaseName 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAddonDefinitionRelationship(t *testing.T) {
	p := models.KubernetesAddonDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAddonDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAddonDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesBaseInfrastructureProviderRelationship(t *testing.T) {
	p := models.KubernetesBaseInfrastructureProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesBaseInfrastructureProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesBaseInfrastructureProviderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesBaseVirtualMachineInfraConfig(t *testing.T) {
	p := models.KubernetesBaseVirtualMachineInfraConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.BaseVirtualMachineInfraConfig","ObjectType":"kubernetes.BaseVirtualMachineInfraConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesBaseVirtualMachineInfraConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesBaseVirtualMachineInfraConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.BaseVirtualMachineInfraConfig", "object_type": "kubernetes.BaseVirtualMachineInfraConfig"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesCatalogRelationship(t *testing.T) {
	p := models.KubernetesCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterRelationship(t *testing.T) {
	p := models.KubernetesClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterAddonProfileRelationship(t *testing.T) {
	p := models.KubernetesClusterAddonProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterAddonProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterAddonProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterCertificateConfiguration(t *testing.T) {
	p := models.KubernetesClusterCertificateConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.ClusterCertificateConfiguration","ObjectType":"kubernetes.ClusterCertificateConfiguration","FrontProxyCert":"FrontProxyCert %d","EtcdCert":"EtcdCert %d","SaPublicKey":"SaPublicKey %d","CaCert":"CaCert %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterCertificateConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterCertificateConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.ClusterCertificateConfiguration", "object_type": "kubernetes.ClusterCertificateConfiguration", "front_proxy_cert": "FrontProxyCert 1", "etcd_cert": "EtcdCert 1", "sa_public_key": "SaPublicKey 1", "ca_cert": "CaCert 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterManagementConfig(t *testing.T) {
	p := models.KubernetesClusterManagementConfig{}
	var d = &schema.ResourceData{}
	c := `{"MasterVip":"MasterVip %d","SshUser":"SshUser %d","LoadBalancerCount":32,"ClassId":"kubernetes.ClusterManagementConfig","ObjectType":"kubernetes.ClusterManagementConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterManagementConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterManagementConfig(p, d)[0]
	expectedOp := map[string]interface{}{"master_vip": "MasterVip 1", "ssh_user": "SshUser 1", "load_balancer_count": 32, "class_id": "kubernetes.ClusterManagementConfig", "object_type": "kubernetes.ClusterManagementConfig"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterProfileRelationship(t *testing.T) {
	p := models.KubernetesClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesCniConfig(t *testing.T) {
	p := models.KubernetesCniConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.CniConfig","ObjectType":"kubernetes.CniConfig","Registry":"Registry %d","Version":"Version %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesCniConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesCniConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.CniConfig", "object_type": "kubernetes.CniConfig", "registry": "Registry 1", "nr_version": "Version 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesConfigResultRelationship(t *testing.T) {
	p := models.KubernetesConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesConfiguration(t *testing.T) {
	p := models.KubernetesConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.Configuration","ObjectType":"kubernetes.Configuration","KubeConfig":"KubeConfig %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.Configuration", "object_type": "kubernetes.Configuration", "kube_config": "KubeConfig 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesContainerRuntimePolicyRelationship(t *testing.T) {
	p := models.KubernetesContainerRuntimePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesContainerRuntimePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesContainerRuntimePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesDaemonSetStatus(t *testing.T) {
	p := models.KubernetesDaemonSetStatus{}
	var d = &schema.ResourceData{}
	c := `{"CurrentNumberScheduled":32,"ObservedGeneration":32,"DesiredNumberScheduled":32,"NumberAvailable":"NumberAvailable %d","NumberMisscheduled":32,"ClassId":"kubernetes.DaemonSetStatus","NumberReady":32,"UpdatedNumberScheduled":"UpdatedNumberScheduled %d","ObjectType":"kubernetes.DaemonSetStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesDaemonSetStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesDaemonSetStatus(p, d)[0]
	expectedOp := map[string]interface{}{"current_number_scheduled": 32, "observed_generation": 32, "desired_number_scheduled": 32, "number_available": "NumberAvailable 1", "number_misscheduled": 32, "class_id": "kubernetes.DaemonSetStatus", "number_ready": 32, "updated_number_scheduled": "UpdatedNumberScheduled 1", "object_type": "kubernetes.DaemonSetStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesDeploymentStatus(t *testing.T) {
	p := models.KubernetesDeploymentStatus{}
	var d = &schema.ResourceData{}
	c := `{"ReadyReplicas":32,"ClassId":"kubernetes.DeploymentStatus","ObjectType":"kubernetes.DeploymentStatus","Replicas":32,"UpdatedReplicas":32,"AvailableReplicas":32,"ObservedGeneration":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesDeploymentStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesDeploymentStatus(p, d)[0]
	expectedOp := map[string]interface{}{"ready_replicas": 32, "class_id": "kubernetes.DeploymentStatus", "object_type": "kubernetes.DeploymentStatus", "replicas": 32, "updated_replicas": 32, "available_replicas": 32, "observed_generation": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesIngressStatus(t *testing.T) {
	p := models.KubernetesIngressStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.IngressStatus","ObjectType":"kubernetes.IngressStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesIngressStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesIngressStatus(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.IngressStatus", "object_type": "kubernetes.IngressStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNetworkPolicyRelationship(t *testing.T) {
	p := models.KubernetesNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeGroupProfileRelationship(t *testing.T) {
	p := models.KubernetesNodeGroupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeGroupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeGroupProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeInfo(t *testing.T) {
	p := models.KubernetesNodeInfo{}
	var d = &schema.ResourceData{}
	c := `{"SystemUuid":"SystemUuid %d","ContainerRuntimeVersion":"ContainerRuntimeVersion %d","KernelVersion":"KernelVersion %d","KubeProxyVersion":"KubeProxyVersion %d","OperatingSystem":"OperatingSystem %d","Architecture":"Architecture %d","ObjectType":"kubernetes.NodeInfo","KubeletVersion":"KubeletVersion %d","OsImage":"OsImage %d","MachineId":"MachineId %d","BootId":"BootId %d","ClassId":"kubernetes.NodeInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeInfo(p, d)[0]
	expectedOp := map[string]interface{}{"system_uuid": "SystemUuid 1", "container_runtime_version": "ContainerRuntimeVersion 1", "kernel_version": "KernelVersion 1", "kube_proxy_version": "KubeProxyVersion 1", "operating_system": "OperatingSystem 1", "architecture": "Architecture 1", "object_type": "kubernetes.NodeInfo", "kubelet_version": "KubeletVersion 1", "os_image": "OsImage 1", "machine_id": "MachineId 1", "boot_id": "BootId 1", "class_id": "kubernetes.NodeInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeProfileRelationship(t *testing.T) {
	p := models.KubernetesNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeSpec(t *testing.T) {
	p := models.KubernetesNodeSpec{}
	var d = &schema.ResourceData{}
	c := `{"PodCidr":"PodCidr %d","ProviderId":"ProviderId %d","ClassId":"kubernetes.NodeSpec","ObjectType":"kubernetes.NodeSpec"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeSpec(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeSpec(p, d)[0]
	expectedOp := map[string]interface{}{"pod_cidr": "PodCidr 1", "provider_id": "ProviderId 1", "class_id": "kubernetes.NodeSpec", "object_type": "kubernetes.NodeSpec"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesObjectMeta(t *testing.T) {
	p := models.KubernetesObjectMeta{}
	var d = &schema.ResourceData{}
	c := `{"Uuid":"Uuid %d","CreationTimestamp":"CreationTimestamp %d","Name":"Name %d","Namespace":"Namespace %d","ClassId":"kubernetes.ObjectMeta","ObjectType":"kubernetes.ObjectMeta","ResourceVersion":"ResourceVersion %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesObjectMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesObjectMeta(p, d)[0]
	expectedOp := map[string]interface{}{"uuid": "Uuid 1", "creation_timestamp": "CreationTimestamp 1", "name": "Name 1", "namespace": "Namespace 1", "class_id": "kubernetes.ObjectMeta", "object_type": "kubernetes.ObjectMeta", "resource_version": "ResourceVersion 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesPodStatus(t *testing.T) {
	p := models.KubernetesPodStatus{}
	var d = &schema.ResourceData{}
	c := `{"QosClass":"QosClass %d","ClassId":"kubernetes.PodStatus","ObjectType":"kubernetes.PodStatus","StartTime":"StartTime %d","HostIp":"HostIp %d","Phase":"Phase %d","PodIp":"PodIp %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesPodStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesPodStatus(p, d)[0]
	expectedOp := map[string]interface{}{"qos_class": "QosClass 1", "class_id": "kubernetes.PodStatus", "object_type": "kubernetes.PodStatus", "start_time": "StartTime 1", "host_ip": "HostIp 1", "phase": "Phase 1", "pod_ip": "PodIp 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesProxyConfig(t *testing.T) {
	p := models.KubernetesProxyConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.ProxyConfig","ObjectType":"kubernetes.ProxyConfig","Username":"Username %d","Hostname":"Hostname %d","IsPasswordSet":true,"Port":32,"Protocol":"Protocol %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesProxyConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesProxyConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.ProxyConfig", "object_type": "kubernetes.ProxyConfig", "username": "Username 1", "hostname": "Hostname 1", "is_password_set": true, "port": 32, "protocol": "Protocol 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesServiceStatus(t *testing.T) {
	p := models.KubernetesServiceStatus{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.ServiceStatus","ClassId":"kubernetes.ServiceStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesServiceStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesServiceStatus(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "kubernetes.ServiceStatus", "class_id": "kubernetes.ServiceStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesStatefulSetStatus(t *testing.T) {
	p := models.KubernetesStatefulSetStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.StatefulSetStatus","ObjectType":"kubernetes.StatefulSetStatus","UpdatedReplicas":32,"AvailableReplicas":32,"ReadyReplicas":32,"CollisionCount":32,"CurrentRevision":"CurrentRevision %d","Replicas":32,"UpdateRevision":"UpdateRevision %d","ObservedGeneration":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesStatefulSetStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesStatefulSetStatus(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.StatefulSetStatus", "object_type": "kubernetes.StatefulSetStatus", "updated_replicas": 32, "available_replicas": 32, "ready_replicas": 32, "collision_count": 32, "current_revision": "CurrentRevision 1", "replicas": 32, "update_revision": "UpdateRevision 1", "observed_generation": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesSysConfigPolicyRelationship(t *testing.T) {
	p := models.KubernetesSysConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesSysConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesSysConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesTrustedRegistriesPolicyRelationship(t *testing.T) {
	p := models.KubernetesTrustedRegistriesPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesTrustedRegistriesPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesTrustedRegistriesPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVersionRelationship(t *testing.T) {
	p := models.KubernetesVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVersionPolicyRelationship(t *testing.T) {
	p := models.KubernetesVersionPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVersionPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVersionPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(t *testing.T) {
	p := models.KubernetesVirtualMachineInfraConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVirtualMachineInstanceTypeRelationship(t *testing.T) {
	p := models.KubernetesVirtualMachineInstanceTypeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVirtualMachineInstanceTypeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVirtualMachineInstanceTypeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKvmSessionRelationship(t *testing.T) {
	p := models.KvmSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKvmSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKvmSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKvmTunnelRelationship(t *testing.T) {
	p := models.KvmTunnelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKvmTunnelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKvmTunnelRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseAccountLicenseDataRelationship(t *testing.T) {
	p := models.LicenseAccountLicenseDataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseAccountLicenseDataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseAccountLicenseDataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseCustomerOpRelationship(t *testing.T) {
	p := models.LicenseCustomerOpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseCustomerOpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseCustomerOpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseIwoCustomerOpRelationship(t *testing.T) {
	p := models.LicenseIwoCustomerOpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseIwoCustomerOpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseIwoCustomerOpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseIwoLicenseCountRelationship(t *testing.T) {
	p := models.LicenseIwoLicenseCountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseIwoLicenseCountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseIwoLicenseCountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseSmartlicenseTokenRelationship(t *testing.T) {
	p := models.LicenseSmartlicenseTokenRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseSmartlicenseTokenRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseSmartlicenseTokenRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolBlock(t *testing.T) {
	p := models.MacpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"macpool.Block","From":"From %d","To":"To %d","ObjectType":"macpool.Block","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolBlock(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "macpool.Block", "from": "From 1", "to": "To 1", "object_type": "macpool.Block", "size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolIdBlockRelationship(t *testing.T) {
	p := models.MacpoolIdBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolIdBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolIdBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolLeaseRelationship(t *testing.T) {
	p := models.MacpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolPoolRelationship(t *testing.T) {
	p := models.MacpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolPoolMemberRelationship(t *testing.T) {
	p := models.MacpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolUniverseRelationship(t *testing.T) {
	p := models.MacpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapManagementControllerRelationship(t *testing.T) {
	p := models.ManagementControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapManagementControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapManagementControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapManagementEntityRelationship(t *testing.T) {
	p := models.ManagementEntityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapManagementEntityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapManagementEntityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryArrayRelationship(t *testing.T) {
	p := models.MemoryArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryConfigResultRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryConfigurationRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryLocalSecurity(t *testing.T) {
	p := models.MemoryPersistentMemoryLocalSecurity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"memory.PersistentMemoryLocalSecurity","ObjectType":"memory.PersistentMemoryLocalSecurity","Enabled":true,"IsSecurePassphraseSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryLocalSecurity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryLocalSecurity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "memory.PersistentMemoryLocalSecurity", "object_type": "memory.PersistentMemoryLocalSecurity", "enabled": true, "is_secure_passphrase_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryRegionRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryRegionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryRegionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryRegionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoBaseMoRelationship(t *testing.T) {
	p := models.MoBaseMoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoBaseMoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoBaseMoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoVersionContext(t *testing.T) {
	p := models.MoVersionContext{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.VersionContext","ClassId":"mo.VersionContext","Version":"Version %d","VersionType":"VersionType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoVersionContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoVersionContext(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.VersionContext", "class_id": "mo.VersionContext", "nr_version": "Version 1", "version_type": "VersionType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkElementRelationship(t *testing.T) {
	p := models.NetworkElementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkElementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkElementRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkFcZoneInfoRelationship(t *testing.T) {
	p := models.NetworkFcZoneInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkFcZoneInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkFcZoneInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkVlanPortInfoRelationship(t *testing.T) {
	p := models.NetworkVlanPortInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkVlanPortInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkVlanPortInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiaapiNewReleaseDetail(t *testing.T) {
	p := models.NiaapiNewReleaseDetail{}
	var d = &schema.ResourceData{}
	c := `{"Version":"Version %d","Description":"Description %d","Link":"Link %d","ReleaseNoteLinkTitle":"ReleaseNoteLinkTitle %d","Title":"Title %d","ClassId":"niaapi.NewReleaseDetail","ObjectType":"niaapi.NewReleaseDetail","ReleaseNoteLink":"ReleaseNoteLink %d","SoftwareDownloadLink":"SoftwareDownloadLink %d","SoftwareDownloadLinkTitle":"SoftwareDownloadLinkTitle %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiaapiNewReleaseDetail(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiaapiNewReleaseDetail(p, d)[0]
	expectedOp := map[string]interface{}{"nr_version": "Version 1", "description": "Description 1", "link": "Link 1", "release_note_link_title": "ReleaseNoteLinkTitle 1", "title": "Title 1", "class_id": "niaapi.NewReleaseDetail", "object_type": "niaapi.NewReleaseDetail", "release_note_link": "ReleaseNoteLink 1", "software_download_link": "SoftwareDownloadLink 1", "software_download_link_title": "SoftwareDownloadLinkTitle 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiaapiVersionRegexPlatform(t *testing.T) {
	p := models.NiaapiVersionRegexPlatform{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niaapi.VersionRegexPlatform","ObjectType":"niaapi.VersionRegexPlatform","Anyllregex":"Anyllregex %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiaapiVersionRegexPlatform(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiaapiVersionRegexPlatform(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niaapi.VersionRegexPlatform", "object_type": "niaapi.VersionRegexPlatform", "anyllregex": "Anyllregex 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryBootflashDetails(t *testing.T) {
	p := models.NiatelemetryBootflashDetails{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.BootflashDetails","ObjectType":"niatelemetry.BootflashDetails","Serial":"Serial %d","FwRev":"FwRev %d","ModelType":"ModelType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryBootflashDetails(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryBootflashDetails(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.BootflashDetails", "object_type": "niatelemetry.BootflashDetails", "serial": "Serial 1", "fw_rev": "FwRev 1", "model_type": "ModelType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryDiskinfo(t *testing.T) {
	p := models.NiatelemetryDiskinfo{}
	var d = &schema.ResourceData{}
	c := `{"Used":32,"Free":32,"Name":"Name %d","ClassId":"niatelemetry.Diskinfo","ObjectType":"niatelemetry.Diskinfo","Total":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryDiskinfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryDiskinfo(p, d)[0]
	expectedOp := map[string]interface{}{"used": 32, "free": 32, "name": "Name 1", "class_id": "niatelemetry.Diskinfo", "object_type": "niatelemetry.Diskinfo", "total": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryInterface(t *testing.T) {
	p := models.NiatelemetryInterface{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.Interface","ObjectType":"niatelemetry.Interface","InterfaceDownCount":32,"InterfaceUpCount":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryInterface(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.Interface", "object_type": "niatelemetry.Interface", "interface_down_count": 32, "interface_up_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNexusDashboardsRelationship(t *testing.T) {
	p := models.NiatelemetryNexusDashboardsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNexusDashboardsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNexusDashboardsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNiaInventoryRelationship(t *testing.T) {
	p := models.NiatelemetryNiaInventoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNiaInventoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNiaInventoryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNiaLicenseStateRelationship(t *testing.T) {
	p := models.NiatelemetryNiaLicenseStateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNiaLicenseStateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNiaLicenseStateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNvePacketCounters(t *testing.T) {
	p := models.NiatelemetryNvePacketCounters{}
	var d = &schema.ResourceData{}
	c := `{"UcastOutpkts":32,"McastInpkts":32,"McastOutbytes":32,"ClassId":"niatelemetry.NvePacketCounters","ObjectType":"niatelemetry.NvePacketCounters","UcastInpkts":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNvePacketCounters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNvePacketCounters(p, d)[0]
	expectedOp := map[string]interface{}{"ucast_outpkts": 32, "mcast_inpkts": 32, "mcast_outbytes": 32, "class_id": "niatelemetry.NvePacketCounters", "object_type": "niatelemetry.NvePacketCounters", "ucast_inpkts": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNveVni(t *testing.T) {
	p := models.NiatelemetryNveVni{}
	var d = &schema.ResourceData{}
	c := `{"DpVniUp":32,"CpVniDown":32,"DpVniDown":32,"ClassId":"niatelemetry.NveVni","ObjectType":"niatelemetry.NveVni","CpVniCount":32,"CpVniUp":32,"DpVniCount":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNveVni(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNveVni(p, d)[0]
	expectedOp := map[string]interface{}{"dp_vni_up": 32, "cp_vni_down": 32, "dp_vni_down": 32, "class_id": "niatelemetry.NveVni", "object_type": "niatelemetry.NveVni", "cp_vni_count": 32, "cp_vni_up": 32, "dp_vni_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNxosBgpMvpn(t *testing.T) {
	p := models.NiatelemetryNxosBgpMvpn{}
	var d = &schema.ResourceData{}
	c := `{"TableVersion":32,"TotalPaths":32,"CapablePeers":32,"ConfiguredPeers":32,"NumberOfClusterLists":32,"ObjectType":"niatelemetry.NxosBgpMvpn","TotalNetworks":32,"ClassId":"niatelemetry.NxosBgpMvpn","MemoryUsed":32,"NumberOfCommunities":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNxosBgpMvpn(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNxosBgpMvpn(p, d)[0]
	expectedOp := map[string]interface{}{"table_version": 32, "total_paths": 32, "capable_peers": 32, "configured_peers": 32, "number_of_cluster_lists": 32, "object_type": "niatelemetry.NxosBgpMvpn", "total_networks": 32, "class_id": "niatelemetry.NxosBgpMvpn", "memory_used": 32, "number_of_communities": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNxosVtp(t *testing.T) {
	p := models.NiatelemetryNxosVtp{}
	var d = &schema.ResourceData{}
	c := `{"Version":32,"ClassId":"niatelemetry.NxosVtp","ObjectType":"niatelemetry.NxosVtp","OperMode":"OperMode %d","PruningMode":"PruningMode %d","TrapEnabled":"TrapEnabled %d","RunningVersion":"RunningVersion %d","V2Mode":"V2Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNxosVtp(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNxosVtp(p, d)[0]
	expectedOp := map[string]interface{}{"nr_version": 32, "class_id": "niatelemetry.NxosVtp", "object_type": "niatelemetry.NxosVtp", "oper_mode": "OperMode 1", "pruning_mode": "PruningMode 1", "trap_enabled": "TrapEnabled 1", "running_version": "RunningVersion 1", "v2_mode": "V2Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetrySmartLicense(t *testing.T) {
	p := models.NiatelemetrySmartLicense{}
	var d = &schema.ResourceData{}
	c := `{"AuthStatus":"AuthStatus %d","LicenseUdi":"LicenseUdi %d","ClassId":"niatelemetry.SmartLicense","ObjectType":"niatelemetry.SmartLicense","SmartAccount":"SmartAccount %d","ActiveMode":"ActiveMode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetrySmartLicense(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetrySmartLicense(p, d)[0]
	expectedOp := map[string]interface{}{"auth_status": "AuthStatus 1", "license_udi": "LicenseUdi 1", "class_id": "niatelemetry.SmartLicense", "object_type": "niatelemetry.SmartLicense", "smart_account": "SmartAccount 1", "active_mode": "ActiveMode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOnpremSchedule(t *testing.T) {
	p := models.OnpremSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.Schedule","DayOfMonth":32,"WeekOfMonth":32,"MonthOfYear":32,"ObjectType":"onprem.Schedule","TimeOfDay":32,"DayOfWeek":32,"RepeatInterval":32,"TimeZone":"TimeZone %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOnpremSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOnpremSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "onprem.Schedule", "day_of_month": 32, "week_of_month": 32, "month_of_year": 32, "object_type": "onprem.Schedule", "time_of_day": 32, "day_of_week": 32, "repeat_interval": 32, "time_zone": "TimeZone 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOnpremUpgradePhase(t *testing.T) {
	p := models.OnpremUpgradePhase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.UpgradePhase","ObjectType":"onprem.UpgradePhase","Message":"Message %d","Name":"Name %d","ElapsedTime":32,"Failed":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapOnpremUpgradePhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOnpremUpgradePhase(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "onprem.UpgradePhase", "object_type": "onprem.UpgradePhase", "message": "Message 1", "name": "Name 1", "elapsed_time": 32, "failed": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOrganizationOrganizationRelationship(t *testing.T) {
	p := models.OrganizationOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOrganizationOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOrganizationOrganizationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsAnswers(t *testing.T) {
	p := models.OsAnswers{}
	var d = &schema.ResourceData{}
	c := `{"IsRootPasswordCrypted":true,"ObjectType":"os.Answers","Hostname":"Hostname %d","ProductKey":"ProductKey %d","IpConfigType":"IpConfigType %d","Nameserver":"Nameserver %d","IsRootPasswordSet":true,"Source":"Source %d","ClassId":"os.Answers","IsAnswerFileSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsAnswers(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsAnswers(p, d)[0]
	expectedOp := map[string]interface{}{"is_root_password_crypted": true, "object_type": "os.Answers", "hostname": "Hostname 1", "product_key": "ProductKey 1", "ip_config_type": "IpConfigType 1", "nameserver": "Nameserver 1", "is_root_password_set": true, "nr_source": "Source 1", "class_id": "os.Answers", "is_answer_file_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsCatalogRelationship(t *testing.T) {
	p := models.OsCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsConfigurationFileRelationship(t *testing.T) {
	p := models.OsConfigurationFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsConfigurationFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsConfigurationFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsGlobalConfig(t *testing.T) {
	p := models.OsGlobalConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.GlobalConfig","OsImageName":"OsImageName %d","ConfigurationSource":"ConfigurationSource %d","InstallMethod":"InstallMethod %d","InstallTargetType":"InstallTargetType %d","ObjectType":"os.GlobalConfig","ScuImageName":"ScuImageName %d","ConfigurationFileName":"ConfigurationFileName %d","WindowsEdition":"WindowsEdition %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsGlobalConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsGlobalConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "os.GlobalConfig", "os_image_name": "OsImageName 1", "configuration_source": "ConfigurationSource 1", "install_method": "InstallMethod 1", "install_target_type": "InstallTargetType 1", "object_type": "os.GlobalConfig", "scu_image_name": "ScuImageName 1", "configuration_file_name": "ConfigurationFileName 1", "windows_edition": "WindowsEdition 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsInstallTarget(t *testing.T) {
	p := models.OsInstallTarget{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.InstallTarget","ObjectType":"os.InstallTarget"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsInstallTarget(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsInstallTarget(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "os.InstallTarget", "object_type": "os.InstallTarget"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsOperatingSystemParameters(t *testing.T) {
	p := models.OsOperatingSystemParameters{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.OperatingSystemParameters","ObjectType":"os.OperatingSystemParameters"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsOperatingSystemParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsOperatingSystemParameters(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "os.OperatingSystemParameters", "object_type": "os.OperatingSystemParameters"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPciSwitchRelationship(t *testing.T) {
	p := models.PciSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPciSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPciSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixDistinguishedName(t *testing.T) {
	p := models.PkixDistinguishedName{}
	var d = &schema.ResourceData{}
	c := `{"CommonName":"CommonName %d","ClassId":"pkix.DistinguishedName","ObjectType":"pkix.DistinguishedName"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixDistinguishedName(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixDistinguishedName(p, d)[0]
	expectedOp := map[string]interface{}{"common_name": "CommonName 1", "class_id": "pkix.DistinguishedName", "object_type": "pkix.DistinguishedName"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixKeyGenerationSpec(t *testing.T) {
	p := models.PkixKeyGenerationSpec{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"pkix.KeyGenerationSpec","ObjectType":"pkix.KeyGenerationSpec","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixKeyGenerationSpec(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixKeyGenerationSpec(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "pkix.KeyGenerationSpec", "object_type": "pkix.KeyGenerationSpec", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixSubjectAlternateName(t *testing.T) {
	p := models.PkixSubjectAlternateName{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"pkix.SubjectAlternateName","ClassId":"pkix.SubjectAlternateName"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixSubjectAlternateName(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixSubjectAlternateName(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "pkix.SubjectAlternateName", "class_id": "pkix.SubjectAlternateName"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyAbstractConfigProfileRelationship(t *testing.T) {
	p := models.PolicyAbstractConfigProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyAbstractConfigProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyAbstractConfigProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyAbstractProfileRelationship(t *testing.T) {
	p := models.PolicyAbstractProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyAbstractProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyAbstractProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigChange(t *testing.T) {
	p := models.PolicyConfigChange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"policy.ConfigChange","ObjectType":"policy.ConfigChange"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigChange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigChange(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "policy.ConfigChange", "object_type": "policy.ConfigChange"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigChangeContext(t *testing.T) {
	p := models.PolicyConfigChangeContext{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"policy.ConfigChangeContext","ObjectType":"policy.ConfigChangeContext","ConfigChangeError":"ConfigChangeError %d","ConfigChangeState":"ConfigChangeState %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigChangeContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigChangeContext(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "policy.ConfigChangeContext", "object_type": "policy.ConfigChangeContext", "config_change_error": "ConfigChangeError 1", "config_change_state": "ConfigChangeState 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigContext(t *testing.T) {
	p := models.PolicyConfigContext{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"policy.ConfigContext","ObjectType":"policy.ConfigContext","ErrorState":"ErrorState %d","OperState":"OperState %d","ConfigState":"ConfigState %d","ControlAction":"ControlAction %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigContext(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "policy.ConfigContext", "object_type": "policy.ConfigContext", "error_state": "ErrorState 1", "oper_state": "OperState 1", "config_state": "ConfigState 1", "control_action": "ControlAction 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigResultContext(t *testing.T) {
	p := models.PolicyConfigResultContext{}
	var d = &schema.ResourceData{}
	c := `{"EntityMoid":"EntityMoid %d","EntityName":"EntityName %d","EntityType":"EntityType %d","ClassId":"policy.ConfigResultContext","ObjectType":"policy.ConfigResultContext"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigResultContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigResultContext(p, d)[0]
	expectedOp := map[string]interface{}{"entity_moid": "EntityMoid 1", "entity_name": "EntityName 1", "entity_type": "EntityType 1", "class_id": "policy.ConfigResultContext", "object_type": "policy.ConfigResultContext"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortGroupRelationship(t *testing.T) {
	p := models.PortGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortInterfaceBaseRelationship(t *testing.T) {
	p := models.PortInterfaceBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortInterfaceBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortInterfaceBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortSubGroupRelationship(t *testing.T) {
	p := models.PortSubGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortSubGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortSubGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPowerControlStateRelationship(t *testing.T) {
	p := models.PowerControlStateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPowerControlStateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPowerControlStateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecommendationCapacityRunwayRelationship(t *testing.T) {
	p := models.RecommendationCapacityRunwayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecommendationCapacityRunwayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecommendationCapacityRunwayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryAbstractBackupInfoRelationship(t *testing.T) {
	p := models.RecoveryAbstractBackupInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryAbstractBackupInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryAbstractBackupInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupConfigPolicyRelationship(t *testing.T) {
	p := models.RecoveryBackupConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupProfileRelationship(t *testing.T) {
	p := models.RecoveryBackupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupSchedule(t *testing.T) {
	p := models.RecoveryBackupSchedule{}
	var d = &schema.ResourceData{}
	c := `{"Hours":32,"ClassId":"recovery.BackupSchedule","ObjectType":"recovery.BackupSchedule","FrequencyUnit":"FrequencyUnit %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"hours": 32, "class_id": "recovery.BackupSchedule", "object_type": "recovery.BackupSchedule", "frequency_unit": "FrequencyUnit 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryConfigParams(t *testing.T) {
	p := models.RecoveryConfigParams{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"recovery.ConfigParams","ObjectType":"recovery.ConfigParams"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryConfigParams(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryConfigParams(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "recovery.ConfigParams", "object_type": "recovery.ConfigParams"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryConfigResultRelationship(t *testing.T) {
	p := models.RecoveryConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryScheduleConfigPolicyRelationship(t *testing.T) {
	p := models.RecoveryScheduleConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryScheduleConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryScheduleConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourceGroupRelationship(t *testing.T) {
	p := models.ResourceGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourceGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourceGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourceMembershipHolderRelationship(t *testing.T) {
	p := models.ResourceMembershipHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourceMembershipHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourceMembershipHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanProfileRelationship(t *testing.T) {
	p := models.SdwanProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanRouterPolicyRelationship(t *testing.T) {
	p := models.SdwanRouterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanRouterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanRouterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanVmanageAccountPolicyRelationship(t *testing.T) {
	p := models.SdwanVmanageAccountPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanVmanageAccountPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanVmanageAccountPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapServerConfigResultRelationship(t *testing.T) {
	p := models.ServerConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapServerConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapServerConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapServerProfileRelationship(t *testing.T) {
	p := models.ServerProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapServerProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapServerProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSessionAbstractSessionRelationship(t *testing.T) {
	p := models.SessionAbstractSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSessionAbstractSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSessionAbstractSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwareHyperflexDistributableRelationship(t *testing.T) {
	p := models.SoftwareHyperflexDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwareHyperflexDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwareHyperflexDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwareSolutionDistributableRelationship(t *testing.T) {
	p := models.SoftwareSolutionDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwareSolutionDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwareSolutionDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryCatalogRelationship(t *testing.T) {
	p := models.SoftwarerepositoryCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryFileRelationship(t *testing.T) {
	p := models.SoftwarerepositoryFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryFileServer(t *testing.T) {
	p := models.SoftwarerepositoryFileServer{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"softwarerepository.FileServer","ObjectType":"softwarerepository.FileServer"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryFileServer(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryFileServer(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "softwarerepository.FileServer", "object_type": "softwarerepository.FileServer"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryOperatingSystemFileRelationship(t *testing.T) {
	p := models.SoftwarerepositoryOperatingSystemFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryReleaseRelationship(t *testing.T) {
	p := models.SoftwarerepositoryReleaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryReleaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryReleaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageBaseCapacity(t *testing.T) {
	p := models.StorageBaseCapacity{}
	var d = &schema.ResourceData{}
	c := `{"Total":32,"Used":32,"Available":32,"CapacityUtilization":32.000000,"Free":32,"ClassId":"storage.BaseCapacity","ObjectType":"storage.BaseCapacity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageBaseCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageBaseCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"total": 32, "used": 32, "available": 32, "capacity_utilization": 32.000000, "free": 32, "class_id": "storage.BaseCapacity", "object_type": "storage.BaseCapacity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageControllerRelationship(t *testing.T) {
	p := models.StorageControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageDiskGroupRelationship(t *testing.T) {
	p := models.StorageDiskGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageDiskGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageDiskGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageDiskSlotRelationship(t *testing.T) {
	p := models.StorageDiskSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageDiskSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageDiskSlotRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageEnclosureRelationship(t *testing.T) {
	p := models.StorageEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageEnclosureRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageFlexFlashControllerRelationship(t *testing.T) {
	p := models.StorageFlexFlashControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageFlexFlashControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageFlexFlashControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageFlexUtilControllerRelationship(t *testing.T) {
	p := models.StorageFlexUtilControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageFlexUtilControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageFlexUtilControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiArrayRelationship(t *testing.T) {
	p := models.StorageHitachiArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiHostRelationship(t *testing.T) {
	p := models.StorageHitachiHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiParityGroupRelationship(t *testing.T) {
	p := models.StorageHitachiParityGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiParityGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiParityGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiPoolRelationship(t *testing.T) {
	p := models.StorageHitachiPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiVolumeRelationship(t *testing.T) {
	p := models.StorageHitachiVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHyperFlexStorageContainerRelationship(t *testing.T) {
	p := models.StorageHyperFlexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHyperFlexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHyperFlexStorageContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppClusterRelationship(t *testing.T) {
	p := models.StorageNetAppClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppEthernetPortRelationship(t *testing.T) {
	p := models.StorageNetAppEthernetPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppEthernetPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppEthernetPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppFcPortRelationship(t *testing.T) {
	p := models.StorageNetAppFcPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppFcPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppFcPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppLunRelationship(t *testing.T) {
	p := models.StorageNetAppLunRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppLunRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppLunRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppNodeRelationship(t *testing.T) {
	p := models.StorageNetAppNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppNodeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppStorageVmRelationship(t *testing.T) {
	p := models.StorageNetAppStorageVmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppStorageVmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppStorageVmRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppVolumeRelationship(t *testing.T) {
	p := models.StorageNetAppVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePhysicalDiskRelationship(t *testing.T) {
	p := models.StoragePhysicalDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePhysicalDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePhysicalDiskRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureArrayRelationship(t *testing.T) {
	p := models.StoragePureArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureControllerRelationship(t *testing.T) {
	p := models.StoragePureControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureHostRelationship(t *testing.T) {
	p := models.StoragePureHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureHostGroupRelationship(t *testing.T) {
	p := models.StoragePureHostGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureHostGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureHostGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureProtectionGroupRelationship(t *testing.T) {
	p := models.StoragePureProtectionGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureProtectionGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureProtectionGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureProtectionGroupSnapshotRelationship(t *testing.T) {
	p := models.StoragePureProtectionGroupSnapshotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureProtectionGroupSnapshotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureProtectionGroupSnapshotRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureVolumeRelationship(t *testing.T) {
	p := models.StoragePureVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageSasExpanderRelationship(t *testing.T) {
	p := models.StorageSasExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageSasExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageSasExpanderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveRelationship(t *testing.T) {
	p := models.StorageVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveContainerRelationship(t *testing.T) {
	p := models.StorageVirtualDriveContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveExtensionRelationship(t *testing.T) {
	p := models.StorageVirtualDriveExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveExtensionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamBaseAdvisoryRelationship(t *testing.T) {
	p := models.TamBaseAdvisoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamBaseAdvisoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamBaseAdvisoryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamBaseAdvisoryDetails(t *testing.T) {
	p := models.TamBaseAdvisoryDetails{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"tam.BaseAdvisoryDetails","ObjectType":"tam.BaseAdvisoryDetails","Description":"Description %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamBaseAdvisoryDetails(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamBaseAdvisoryDetails(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "tam.BaseAdvisoryDetails", "object_type": "tam.BaseAdvisoryDetails", "description": "Description 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamSeverity(t *testing.T) {
	p := models.TamSeverity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"tam.Severity","ObjectType":"tam.Severity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamSeverity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamSeverity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "tam.Severity", "object_type": "tam.Severity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTechsupportmanagementTechSupportBundleRelationship(t *testing.T) {
	p := models.TechsupportmanagementTechSupportBundleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTechsupportmanagementTechSupportBundleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTechsupportmanagementTechSupportBundleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTechsupportmanagementTechSupportStatusRelationship(t *testing.T) {
	p := models.TechsupportmanagementTechSupportStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTechsupportmanagementTechSupportStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTechsupportmanagementTechSupportStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTopSystemRelationship(t *testing.T) {
	p := models.TopSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTopSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTopSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolBlockRelationship(t *testing.T) {
	p := models.UuidpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolPoolRelationship(t *testing.T) {
	p := models.UuidpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolPoolMemberRelationship(t *testing.T) {
	p := models.UuidpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUniverseRelationship(t *testing.T) {
	p := models.UuidpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUuidBlock(t *testing.T) {
	p := models.UuidpoolUuidBlock{}
	var d = &schema.ResourceData{}
	c := `{"To":"To %d","From":"From %d","Size":32,"ClassId":"uuidpool.UuidBlock","ObjectType":"uuidpool.UuidBlock"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUuidBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUuidBlock(p, d)[0]
	expectedOp := map[string]interface{}{"to": "To 1", "from": "From 1", "size": 32, "class_id": "uuidpool.UuidBlock", "object_type": "uuidpool.UuidBlock"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUuidLeaseRelationship(t *testing.T) {
	p := models.UuidpoolUuidLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUuidLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUuidLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationActionInfo(t *testing.T) {
	p := models.VirtualizationActionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.ActionInfo","ObjectType":"virtualization.ActionInfo","FailureReason":"FailureReason %d","Name":"Name %d","Status":"Status %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationActionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationActionInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.ActionInfo", "object_type": "virtualization.ActionInfo", "failure_reason": "FailureReason 1", "name": "Name 1", "status": "Status 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseClusterRelationship(t *testing.T) {
	p := models.VirtualizationBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseHostRelationship(t *testing.T) {
	p := models.VirtualizationBaseHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseNetworkRelationship(t *testing.T) {
	p := models.VirtualizationBaseNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVirtualDiskRelationship(t *testing.T) {
	p := models.VirtualizationBaseVirtualDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVirtualDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVirtualDiskRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationBaseVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVmConfiguration(t *testing.T) {
	p := models.VirtualizationBaseVmConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.BaseVmConfiguration","ObjectType":"virtualization.BaseVmConfiguration"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVmConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVmConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.BaseVmConfiguration", "object_type": "virtualization.BaseVmConfiguration"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCloudInitConfig(t *testing.T) {
	p := models.VirtualizationCloudInitConfig{}
	var d = &schema.ResourceData{}
	c := `{"NetworkData":"NetworkData %d","NetworkDataBase64Encoded":true,"ClassId":"virtualization.CloudInitConfig","ObjectType":"virtualization.CloudInitConfig","UserData":"UserData %d","UserDataBase64Encoded":true,"ConfigType":"ConfigType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCloudInitConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCloudInitConfig(p, d)[0]
	expectedOp := map[string]interface{}{"network_data": "NetworkData 1", "network_data_base64_encoded": true, "class_id": "virtualization.CloudInitConfig", "object_type": "virtualization.CloudInitConfig", "user_data": "UserData 1", "user_data_base64_encoded": true, "config_type": "ConfigType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationComputeCapacity(t *testing.T) {
	p := models.VirtualizationComputeCapacity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.ComputeCapacity","ObjectType":"virtualization.ComputeCapacity","Free":32,"Used":32,"Capacity":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationComputeCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationComputeCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.ComputeCapacity", "object_type": "virtualization.ComputeCapacity", "free": 32, "used": 32, "capacity": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCpuAllocation(t *testing.T) {
	p := models.VirtualizationCpuAllocation{}
	var d = &schema.ResourceData{}
	c := `{"Used":32,"Free":32,"Reserved":32,"Total":32,"ClassId":"virtualization.CpuAllocation","ObjectType":"virtualization.CpuAllocation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCpuAllocation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCpuAllocation(p, d)[0]
	expectedOp := map[string]interface{}{"used": 32, "free": 32, "reserved": 32, "total": 32, "class_id": "virtualization.CpuAllocation", "object_type": "virtualization.CpuAllocation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCpuInfo(t *testing.T) {
	p := models.VirtualizationCpuInfo{}
	var d = &schema.ResourceData{}
	c := `{"Description":"Description %d","Sockets":32,"Speed":32,"Vendor":"Vendor %d","Cores":32,"ClassId":"virtualization.CpuInfo","ObjectType":"virtualization.CpuInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCpuInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCpuInfo(p, d)[0]
	expectedOp := map[string]interface{}{"description": "Description 1", "sockets": 32, "speed": 32, "vendor": "Vendor 1", "cores": 32, "class_id": "virtualization.CpuInfo", "object_type": "virtualization.CpuInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationGuestInfo(t *testing.T) {
	p := models.VirtualizationGuestInfo{}
	var d = &schema.ResourceData{}
	c := `{"OperatingSystem":"OperatingSystem %d","ClassId":"virtualization.GuestInfo","ObjectType":"virtualization.GuestInfo","Hostname":"Hostname %d","IpAddress":"IpAddress %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationGuestInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationGuestInfo(p, d)[0]
	expectedOp := map[string]interface{}{"operating_system": "OperatingSystem 1", "class_id": "virtualization.GuestInfo", "object_type": "virtualization.GuestInfo", "hostname": "Hostname 1", "ip_address": "IpAddress 1", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationMemoryAllocation(t *testing.T) {
	p := models.VirtualizationMemoryAllocation{}
	var d = &schema.ResourceData{}
	c := `{"Reserved":32,"Total":32,"Used":32,"ClassId":"virtualization.MemoryAllocation","ObjectType":"virtualization.MemoryAllocation","Free":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationMemoryAllocation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationMemoryAllocation(p, d)[0]
	expectedOp := map[string]interface{}{"reserved": 32, "total": 32, "used": 32, "class_id": "virtualization.MemoryAllocation", "object_type": "virtualization.MemoryAllocation", "free": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationMemoryCapacity(t *testing.T) {
	p := models.VirtualizationMemoryCapacity{}
	var d = &schema.ResourceData{}
	c := `{"Free":32,"Used":32,"ClassId":"virtualization.MemoryCapacity","ObjectType":"virtualization.MemoryCapacity","Capacity":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationMemoryCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationMemoryCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"free": 32, "used": 32, "class_id": "virtualization.MemoryCapacity", "object_type": "virtualization.MemoryCapacity", "capacity": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationProductInfo(t *testing.T) {
	p := models.VirtualizationProductInfo{}
	var d = &schema.ResourceData{}
	c := `{"Build":"Build %d","ObjectType":"virtualization.ProductInfo","ClassId":"virtualization.ProductInfo","ProductName":"ProductName %d","ProductType":"ProductType %d","ProductVendor":"ProductVendor %d","Version":"Version %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationProductInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationProductInfo(p, d)[0]
	expectedOp := map[string]interface{}{"build": "Build 1", "object_type": "virtualization.ProductInfo", "class_id": "virtualization.ProductInfo", "product_name": "ProductName 1", "product_type": "ProductType 1", "product_vendor": "ProductVendor 1", "nr_version": "Version 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationStorageCapacity(t *testing.T) {
	p := models.VirtualizationStorageCapacity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.StorageCapacity","ObjectType":"virtualization.StorageCapacity","Capacity":32,"Free":32,"Used":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationStorageCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationStorageCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.StorageCapacity", "object_type": "virtualization.StorageCapacity", "capacity": 32, "free": 32, "used": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareClusterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDatacenterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDatacenterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDatacenterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDatacenterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDatastoreRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDatastoreRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDatastoreRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDatastoreRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDistributedNetworkRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDistributedNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDistributedNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDistributedNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDistributedSwitchRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDistributedSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDistributedSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDistributedSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareFolderRelationship(t *testing.T) {
	p := models.VirtualizationVmwareFolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareFolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareFolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareHostRelationship(t *testing.T) {
	p := models.VirtualizationVmwareHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareNetworkRelationship(t *testing.T) {
	p := models.VirtualizationVmwareNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(t *testing.T) {
	p := models.VirtualizationVmwarePhysicalNetworkInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareRemoteDisplayInfo(t *testing.T) {
	p := models.VirtualizationVmwareRemoteDisplayInfo{}
	var d = &schema.ResourceData{}
	c := `{"RemoteDisplayVncPort":32,"RemoteDisplayPassword":"RemoteDisplayPassword %d","RemoteDisplayVncKey":"RemoteDisplayVncKey %d","ClassId":"virtualization.VmwareRemoteDisplayInfo","ObjectType":"virtualization.VmwareRemoteDisplayInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareRemoteDisplayInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareRemoteDisplayInfo(p, d)[0]
	expectedOp := map[string]interface{}{"remote_display_vnc_port": 32, "remote_display_password": "RemoteDisplayPassword 1", "remote_display_vnc_key": "RemoteDisplayVncKey 1", "class_id": "virtualization.VmwareRemoteDisplayInfo", "object_type": "virtualization.VmwareRemoteDisplayInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareResourceConsumption(t *testing.T) {
	p := models.VirtualizationVmwareResourceConsumption{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.VmwareResourceConsumption","ClassId":"virtualization.VmwareResourceConsumption","CpuConsumed":32,"MemoryConsumed":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareResourceConsumption(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareResourceConsumption(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "virtualization.VmwareResourceConsumption", "class_id": "virtualization.VmwareResourceConsumption", "cpu_consumed": 32, "memory_consumed": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareSharesInfo(t *testing.T) {
	p := models.VirtualizationVmwareSharesInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareSharesInfo","ObjectType":"virtualization.VmwareSharesInfo","Level":"Level %d","Shares":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareSharesInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareSharesInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareSharesInfo", "object_type": "virtualization.VmwareSharesInfo", "level": "Level 1", "shares": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareTeamingAndFailover(t *testing.T) {
	p := models.VirtualizationVmwareTeamingAndFailover{}
	var d = &schema.ResourceData{}
	c := `{"Failback":true,"NetworkFailureDetection":"NetworkFailureDetection %d","LoadBalancing":"LoadBalancing %d","Name":"Name %d","NotifySwitches":true,"ClassId":"virtualization.VmwareTeamingAndFailover","ObjectType":"virtualization.VmwareTeamingAndFailover"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareTeamingAndFailover(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareTeamingAndFailover(p, d)[0]
	expectedOp := map[string]interface{}{"failback": true, "network_failure_detection": "NetworkFailureDetection 1", "load_balancing": "LoadBalancing 1", "name": "Name 1", "notify_switches": true, "class_id": "virtualization.VmwareTeamingAndFailover", "object_type": "virtualization.VmwareTeamingAndFailover"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVcenterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVcenterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVcenterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVcenterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVirtualSwitchRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVirtualSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVirtualSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVirtualSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmCpuShareInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmCpuShareInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareVmCpuShareInfo","ObjectType":"virtualization.VmwareVmCpuShareInfo","CpuShares":32,"CpuLimit":32,"CpuOverheadLimit":32,"CpuReservation":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmCpuShareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmCpuShareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareVmCpuShareInfo", "object_type": "virtualization.VmwareVmCpuShareInfo", "cpu_shares": 32, "cpu_limit": 32, "cpu_overhead_limit": 32, "cpu_reservation": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmCpuSocketInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmCpuSocketInfo{}
	var d = &schema.ResourceData{}
	c := `{"NumSockets":32,"ClassId":"virtualization.VmwareVmCpuSocketInfo","ObjectType":"virtualization.VmwareVmCpuSocketInfo","CoresPerSocket":32,"NumCpus":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmCpuSocketInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmCpuSocketInfo(p, d)[0]
	expectedOp := map[string]interface{}{"num_sockets": 32, "class_id": "virtualization.VmwareVmCpuSocketInfo", "object_type": "virtualization.VmwareVmCpuSocketInfo", "cores_per_socket": 32, "num_cpus": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmDiskCommitInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmDiskCommitInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareVmDiskCommitInfo","ObjectType":"virtualization.VmwareVmDiskCommitInfo","CommittedDisk":32,"UnCommittedDisk":32,"UnsharedDisk":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmDiskCommitInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmDiskCommitInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareVmDiskCommitInfo", "object_type": "virtualization.VmwareVmDiskCommitInfo", "committed_disk": 32, "un_committed_disk": 32, "unshared_disk": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmMemoryShareInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmMemoryShareInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareVmMemoryShareInfo","ObjectType":"virtualization.VmwareVmMemoryShareInfo","MemLimit":32,"MemOverheadLimit":32,"MemReservation":32,"MemShares":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmMemoryShareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmMemoryShareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareVmMemoryShareInfo", "object_type": "virtualization.VmwareVmMemoryShareInfo", "mem_limit": 32, "mem_overhead_limit": 32, "mem_reservation": 32, "mem_shares": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicArfsSettings(t *testing.T) {
	p := models.VnicArfsSettings{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"ClassId":"vnic.ArfsSettings","ObjectType":"vnic.ArfsSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicArfsSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicArfsSettings(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "class_id": "vnic.ArfsSettings", "object_type": "vnic.ArfsSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicCdn(t *testing.T) {
	p := models.VnicCdn{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.Cdn","ObjectType":"vnic.Cdn","Value":"Value %d","Source":"Source %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicCdn(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicCdn(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.Cdn", "object_type": "vnic.Cdn", "value": "Value 1", "nr_source": "Source 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicCompletionQueueSettings(t *testing.T) {
	p := models.VnicCompletionQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"vnic.CompletionQueueSettings","ClassId":"vnic.CompletionQueueSettings","Count":32,"RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicCompletionQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicCompletionQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "vnic.CompletionQueueSettings", "class_id": "vnic.CompletionQueueSettings", "nr_count": 32, "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicEthAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthIfRelationship(t *testing.T) {
	p := models.VnicEthIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthIfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthInterruptSettings(t *testing.T) {
	p := models.VnicEthInterruptSettings{}
	var d = &schema.ResourceData{}
	c := `{"Count":32,"Mode":"Mode %d","CoalescingTime":32,"ClassId":"vnic.EthInterruptSettings","ObjectType":"vnic.EthInterruptSettings","CoalescingType":"CoalescingType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthInterruptSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthInterruptSettings(p, d)[0]
	expectedOp := map[string]interface{}{"nr_count": 32, "mode": "Mode 1", "coalescing_time": 32, "class_id": "vnic.EthInterruptSettings", "object_type": "vnic.EthInterruptSettings", "coalescing_type": "CoalescingType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthNetworkPolicyRelationship(t *testing.T) {
	p := models.VnicEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthQosPolicyRelationship(t *testing.T) {
	p := models.VnicEthQosPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthQosPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthQosPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthRxQueueSettings(t *testing.T) {
	p := models.VnicEthRxQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"RingSize":32,"Count":32,"ClassId":"vnic.EthRxQueueSettings","ObjectType":"vnic.EthRxQueueSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthRxQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthRxQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"ring_size": 32, "nr_count": 32, "class_id": "vnic.EthRxQueueSettings", "object_type": "vnic.EthRxQueueSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthTxQueueSettings(t *testing.T) {
	p := models.VnicEthTxQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"Count":32,"RingSize":32,"ClassId":"vnic.EthTxQueueSettings","ObjectType":"vnic.EthTxQueueSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthTxQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthTxQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"nr_count": 32, "ring_size": 32, "class_id": "vnic.EthTxQueueSettings", "object_type": "vnic.EthTxQueueSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicFcAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcErrorRecoverySettings(t *testing.T) {
	p := models.VnicFcErrorRecoverySettings{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"IoRetryCount":32,"IoRetryTimeout":32,"LinkDownTimeout":32,"PortDownTimeout":32,"ObjectType":"vnic.FcErrorRecoverySettings","ClassId":"vnic.FcErrorRecoverySettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcErrorRecoverySettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcErrorRecoverySettings(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "io_retry_count": 32, "io_retry_timeout": 32, "link_down_timeout": 32, "port_down_timeout": 32, "object_type": "vnic.FcErrorRecoverySettings", "class_id": "vnic.FcErrorRecoverySettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcIfRelationship(t *testing.T) {
	p := models.VnicFcIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcIfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcInterruptSettings(t *testing.T) {
	p := models.VnicFcInterruptSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FcInterruptSettings","ObjectType":"vnic.FcInterruptSettings","Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcInterruptSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcInterruptSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FcInterruptSettings", "object_type": "vnic.FcInterruptSettings", "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcNetworkPolicyRelationship(t *testing.T) {
	p := models.VnicFcNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcQosPolicyRelationship(t *testing.T) {
	p := models.VnicFcQosPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcQosPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcQosPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcQueueSettings(t *testing.T) {
	p := models.VnicFcQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FcQueueSettings","ObjectType":"vnic.FcQueueSettings","Count":32,"RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FcQueueSettings", "object_type": "vnic.FcQueueSettings", "nr_count": 32, "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFlogiSettings(t *testing.T) {
	p := models.VnicFlogiSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FlogiSettings","ObjectType":"vnic.FlogiSettings","Retries":32,"Timeout":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFlogiSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFlogiSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FlogiSettings", "object_type": "vnic.FlogiSettings", "retries": 32, "timeout": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiAuthProfile(t *testing.T) {
	p := models.VnicIscsiAuthProfile{}
	var d = &schema.ResourceData{}
	c := `{"IsPasswordSet":true,"UserId":"UserId %d","ClassId":"vnic.IscsiAuthProfile","ObjectType":"vnic.IscsiAuthProfile"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiAuthProfile(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiAuthProfile(p, d)[0]
	expectedOp := map[string]interface{}{"is_password_set": true, "user_id": "UserId 1", "class_id": "vnic.IscsiAuthProfile", "object_type": "vnic.IscsiAuthProfile"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiBootPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiBootPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiBootPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiBootPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiStaticTargetPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiStaticTargetPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiStaticTargetPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiStaticTargetPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicLanConnectivityPolicyRelationship(t *testing.T) {
	p := models.VnicLanConnectivityPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicLanConnectivityPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicLanConnectivityPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicLun(t *testing.T) {
	p := models.VnicLun{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.Lun","ObjectType":"vnic.Lun","Bootable":true,"LunId":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicLun(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicLun(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.Lun", "object_type": "vnic.Lun", "bootable": true, "lun_id": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicNvgreSettings(t *testing.T) {
	p := models.VnicNvgreSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.NvgreSettings","ObjectType":"vnic.NvgreSettings","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicNvgreSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicNvgreSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.NvgreSettings", "object_type": "vnic.NvgreSettings", "enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicPlacementSettings(t *testing.T) {
	p := models.VnicPlacementSettings{}
	var d = &schema.ResourceData{}
	c := `{"PciLink":32,"SwitchId":"SwitchId %d","Uplink":32,"ClassId":"vnic.PlacementSettings","ObjectType":"vnic.PlacementSettings","Id":"Id %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicPlacementSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicPlacementSettings(p, d)[0]
	expectedOp := map[string]interface{}{"pci_link": 32, "switch_id": "SwitchId 1", "uplink": 32, "class_id": "vnic.PlacementSettings", "object_type": "vnic.PlacementSettings", "id": "Id 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicPlogiSettings(t *testing.T) {
	p := models.VnicPlogiSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.PlogiSettings","ObjectType":"vnic.PlogiSettings","Retries":32,"Timeout":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicPlogiSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicPlogiSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.PlogiSettings", "object_type": "vnic.PlogiSettings", "retries": 32, "timeout": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicRoceSettings(t *testing.T) {
	p := models.VnicRoceSettings{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"vnic.RoceSettings","MemoryRegions":32,"Version":32,"ResourceGroups":32,"ClassId":"vnic.RoceSettings","ClassOfService":32,"Enabled":true,"QueuePairs":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicRoceSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicRoceSettings(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "vnic.RoceSettings", "memory_regions": 32, "nr_version": 32, "resource_groups": 32, "class_id": "vnic.RoceSettings", "class_of_service": 32, "enabled": true, "queue_pairs": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicRssHashSettings(t *testing.T) {
	p := models.VnicRssHashSettings{}
	var d = &schema.ResourceData{}
	c := `{"Ipv6ExtHash":true,"TcpIpv6Hash":true,"ObjectType":"vnic.RssHashSettings","UdpIpv6Hash":true,"TcpIpv6ExtHash":true,"Ipv4Hash":true,"Ipv6Hash":true,"TcpIpv4Hash":true,"UdpIpv4Hash":true,"ClassId":"vnic.RssHashSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicRssHashSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicRssHashSettings(p, d)[0]
	expectedOp := map[string]interface{}{"ipv6_ext_hash": true, "tcp_ipv6_hash": true, "object_type": "vnic.RssHashSettings", "udp_ipv6_hash": true, "tcp_ipv6_ext_hash": true, "ipv4_hash": true, "ipv6_hash": true, "tcp_ipv4_hash": true, "udp_ipv4_hash": true, "class_id": "vnic.RssHashSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicSanConnectivityPolicyRelationship(t *testing.T) {
	p := models.VnicSanConnectivityPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicSanConnectivityPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicSanConnectivityPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicScsiQueueSettings(t *testing.T) {
	p := models.VnicScsiQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.ScsiQueueSettings","ObjectType":"vnic.ScsiQueueSettings","Count":32,"RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicScsiQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicScsiQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.ScsiQueueSettings", "object_type": "vnic.ScsiQueueSettings", "nr_count": 32, "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicTcpOffloadSettings(t *testing.T) {
	p := models.VnicTcpOffloadSettings{}
	var d = &schema.ResourceData{}
	c := `{"LargeReceive":true,"ClassId":"vnic.TcpOffloadSettings","ObjectType":"vnic.TcpOffloadSettings","LargeSend":true,"RxChecksum":true,"TxChecksum":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicTcpOffloadSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicTcpOffloadSettings(p, d)[0]
	expectedOp := map[string]interface{}{"large_receive": true, "class_id": "vnic.TcpOffloadSettings", "object_type": "vnic.TcpOffloadSettings", "large_send": true, "rx_checksum": true, "tx_checksum": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicUsnicSettings(t *testing.T) {
	p := models.VnicUsnicSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.UsnicSettings","ObjectType":"vnic.UsnicSettings","UsnicAdapterPolicy":"UsnicAdapterPolicy %d","Cos":32,"Count":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicUsnicSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicUsnicSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.UsnicSettings", "object_type": "vnic.UsnicSettings", "usnic_adapter_policy": "UsnicAdapterPolicy 1", "cos": 32, "nr_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVlanSettings(t *testing.T) {
	p := models.VnicVlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"AllowedVlans":"AllowedVlans %d","DefaultVlan":32,"Mode":"Mode %d","ClassId":"vnic.VlanSettings","ObjectType":"vnic.VlanSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"allowed_vlans": "AllowedVlans 1", "default_vlan": 32, "mode": "Mode 1", "class_id": "vnic.VlanSettings", "object_type": "vnic.VlanSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVmqSettings(t *testing.T) {
	p := models.VnicVmqSettings{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"MultiQueueSupport":true,"NumInterrupts":32,"NumSubVnics":32,"NumVmqs":32,"ClassId":"vnic.VmqSettings","ObjectType":"vnic.VmqSettings","VmmqAdapterPolicy":"VmmqAdapterPolicy %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVmqSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVmqSettings(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "multi_queue_support": true, "num_interrupts": 32, "num_sub_vnics": 32, "num_vmqs": 32, "class_id": "vnic.VmqSettings", "object_type": "vnic.VmqSettings", "vmmq_adapter_policy": "VmmqAdapterPolicy 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVsanSettings(t *testing.T) {
	p := models.VnicVsanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"vnic.VsanSettings","ClassId":"vnic.VsanSettings","DefaultVlanId":32,"Id":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVsanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVsanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "vnic.VsanSettings", "class_id": "vnic.VsanSettings", "default_vlan_id": 32, "id": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVxlanSettings(t *testing.T) {
	p := models.VnicVxlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.VxlanSettings","ObjectType":"vnic.VxlanSettings","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVxlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVxlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.VxlanSettings", "object_type": "vnic.VxlanSettings", "enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVrfVrfRelationship(t *testing.T) {
	p := models.VrfVrfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVrfVrfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVrfVrfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCatalogRelationship(t *testing.T) {
	p := models.WorkflowCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowComments(t *testing.T) {
	p := models.WorkflowComments{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.Comments","ObjectType":"workflow.Comments","Description":"Description %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowComments(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowComments(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.Comments", "object_type": "workflow.Comments", "description": "Description 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCustomDataTypeDefinitionRelationship(t *testing.T) {
	p := models.WorkflowCustomDataTypeDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCustomDataTypeDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCustomDataTypeDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCustomDataTypeProperties(t *testing.T) {
	p := models.WorkflowCustomDataTypeProperties{}
	var d = &schema.ResourceData{}
	c := `{"Cloneable":true,"ExternalMeta":true,"ClassId":"workflow.CustomDataTypeProperties","ObjectType":"workflow.CustomDataTypeProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCustomDataTypeProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCustomDataTypeProperties(p, d)[0]
	expectedOp := map[string]interface{}{"cloneable": true, "external_meta": true, "class_id": "workflow.CustomDataTypeProperties", "object_type": "workflow.CustomDataTypeProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowErrorResponseHandlerRelationship(t *testing.T) {
	p := models.WorkflowErrorResponseHandlerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowErrorResponseHandlerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowErrorResponseHandlerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowInternalProperties(t *testing.T) {
	p := models.WorkflowInternalProperties{}
	var d = &schema.ResourceData{}
	c := `{"BaseTaskType":"BaseTaskType %d","ObjectType":"workflow.InternalProperties","ClassId":"workflow.InternalProperties","Internal":true,"Owner":"Owner %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowInternalProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowInternalProperties(p, d)[0]
	expectedOp := map[string]interface{}{"base_task_type": "BaseTaskType 1", "object_type": "workflow.InternalProperties", "class_id": "workflow.InternalProperties", "internal": true, "owner": "Owner 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowPendingDynamicWorkflowInfoRelationship(t *testing.T) {
	p := models.WorkflowPendingDynamicWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowProperties(t *testing.T) {
	p := models.WorkflowProperties{}
	var d = &schema.ResourceData{}
	c := `{"RetryDelay":32,"Timeout":32,"ExternalMeta":true,"RetryPolicy":"RetryPolicy %d","RetryCount":32,"TimeoutPolicy":"TimeoutPolicy %d","SupportStatus":"SupportStatus %d","ObjectType":"workflow.Properties","ClassId":"workflow.Properties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowProperties(p, d)[0]
	expectedOp := map[string]interface{}{"retry_delay": 32, "timeout": 32, "external_meta": true, "retry_policy": "RetryPolicy 1", "retry_count": 32, "timeout_policy": "TimeoutPolicy 1", "support_status": "SupportStatus 1", "object_type": "workflow.Properties", "class_id": "workflow.Properties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskConstraints(t *testing.T) {
	p := models.WorkflowTaskConstraints{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.TaskConstraints","ObjectType":"workflow.TaskConstraints"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskConstraints(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskConstraints(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.TaskConstraints", "object_type": "workflow.TaskConstraints"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskDefinitionRelationship(t *testing.T) {
	p := models.WorkflowTaskDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskInfoRelationship(t *testing.T) {
	p := models.WorkflowTaskInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskMetadataRelationship(t *testing.T) {
	p := models.WorkflowTaskMetadataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskMetadataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskMetadataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowValidationInformation(t *testing.T) {
	p := models.WorkflowValidationInformation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.ValidationInformation","ObjectType":"workflow.ValidationInformation","State":"State %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowValidationInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowValidationInformation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.ValidationInformation", "object_type": "workflow.ValidationInformation", "state": "State 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowCtx(t *testing.T) {
	p := models.WorkflowWorkflowCtx{}
	var d = &schema.ResourceData{}
	c := `{"WorkflowMetaName":"WorkflowMetaName %d","WorkflowSubtype":"WorkflowSubtype %d","WorkflowType":"WorkflowType %d","ClassId":"workflow.WorkflowCtx","ObjectType":"workflow.WorkflowCtx"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowCtx(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowCtx(p, d)[0]
	expectedOp := map[string]interface{}{"workflow_meta_name": "WorkflowMetaName 1", "workflow_subtype": "WorkflowSubtype 1", "workflow_type": "WorkflowType 1", "class_id": "workflow.WorkflowCtx", "object_type": "workflow.WorkflowCtx"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowDefinitionRelationship(t *testing.T) {
	p := models.WorkflowWorkflowDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowInfoRelationship(t *testing.T) {
	p := models.WorkflowWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowInfoProperties(t *testing.T) {
	p := models.WorkflowWorkflowInfoProperties{}
	var d = &schema.ResourceData{}
	c := `{"RollbackAction":"RollbackAction %d","ObjectType":"workflow.WorkflowInfoProperties","ClassId":"workflow.WorkflowInfoProperties","Retryable":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowInfoProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowInfoProperties(p, d)[0]
	expectedOp := map[string]interface{}{"rollback_action": "RollbackAction 1", "object_type": "workflow.WorkflowInfoProperties", "class_id": "workflow.WorkflowInfoProperties", "retryable": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowMetadataRelationship(t *testing.T) {
	p := models.WorkflowWorkflowMetadataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowMetadataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowMetadataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowProperties(t *testing.T) {
	p := models.WorkflowWorkflowProperties{}
	var d = &schema.ResourceData{}
	c := `{"EnableDebug":true,"ExternalMeta":true,"Retryable":true,"SupportStatus":"SupportStatus %d","ObjectType":"workflow.WorkflowProperties","ClassId":"workflow.WorkflowProperties","Cloneable":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowProperties(p, d)[0]
	expectedOp := map[string]interface{}{"enable_debug": true, "external_meta": true, "retryable": true, "support_status": "SupportStatus 1", "object_type": "workflow.WorkflowProperties", "class_id": "workflow.WorkflowProperties", "cloneable": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapX509Certificate(t *testing.T) {
	p := models.X509Certificate{}
	var d = &schema.ResourceData{}
	c := `{"PemCertificate":"PemCertificate %d","Sha256Fingerprint":"Sha256Fingerprint %d","SignatureAlgorithm":"SignatureAlgorithm %d","ClassId":"x509.Certificate","ObjectType":"x509.Certificate"}`

	//test when the response is empty
	ffOpEmpty := flattenMapX509Certificate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapX509Certificate(p, d)[0]
	expectedOp := map[string]interface{}{"pem_certificate": "PemCertificate 1", "sha256_fingerprint": "Sha256Fingerprint 1", "signature_algorithm": "SignatureAlgorithm 1", "class_id": "x509.Certificate", "object_type": "x509.Certificate"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}