---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_https_acl_filter_details"
description: |-
        Object to capture the HTTPS ACL EPGs filter details in APIC.

---

# Data Source: intersight_niatelemetry_https_acl_filter_details
Object to capture the HTTPS ACL EPGs filter details in APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_https_acl_filter_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dest_from_port`:(string) Destination From Port HTTPS ACL EPGs filter MO for APIC. 
* `dest_to_port`:(string) Destination To Port HTTPS ACL EPGs filter MO for APIC. 
* `dn`:(string) Dn of the HTTPS ACL EPGs filter MO for APIC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filter_name`:(string) Name of HTTPS ACL filter for APIC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `prot`:(string) Prot of the HTTPS ACL EPGs filter MO for APIC. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
* `src_from_port`:(string) Source From Port HTTPS ACL EPGs filter MO for APIC. 
* `src_to_port`:(string) Source To Port HTTPS ACL EPGs filter MO for APIC. 
 
