package ldap

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/l-with/terraform-provider-ldap/client"
)

func dataSourceLDAPEntry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLDAPEntryRead,
		Schema: map[string]*schema.Schema{
			"dn": {
				Description: "DN of the LDAP entry",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ou": {
				Description: "OU where LDAP entry will be searched",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter": {
				Description: "filter for selecting the LDAP entry",
				Type:        schema.TypeString,
				Required:    true,
			},
			"data_json": {
				Description: "JSON-encoded string that is read as the values of the attributes of the entry (s. https://pkg.go.dev/github.com/go-ldap/ldap/v3#EntryAttribute)",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ignore_attributes": {
				Description: "list of attributes to ignore",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ignore_attribute_patterns": {
				Description: "list of attribute patterns to ignore",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"base64encode_attributes": {
				Description: "list of attributes to be encoded to base64",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"base64encode_attribute_patterns": {
				Description: "list of attribute patterns to be encoded to base64",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceLDAPEntryRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cl := m.(*client.Client)

	ou := d.Get("ou").(string)
	filter := d.Get("filter").(string)
	var ignore_attributes []string
	for _, ignore_attribute := range d.Get("ignore_attributes").([]interface{}) {
		ignore_attributes = append(ignore_attributes, ignore_attribute.(string))
	}
	var ignore_attribute_patterns []string
	for _, ignore_attribute_pattern := range d.Get("ignore_attribute_patterns").([]interface{}) {
		ignore_attribute_patterns = append(ignore_attribute_patterns, ignore_attribute_pattern.(string))
	}
	var base64encode_attributes []string
	for _, base64encode_attribute := range d.Get("base64encode_attributes").([]interface{}) {
		base64encode_attributes = append(base64encode_attributes, base64encode_attribute.(string))
	}
	var base64encode_attribute_patterns []string
	for _, base64encode_attribute_pattern := range d.Get("base64encode_attribute_patterns").([]interface{}) {
		base64encode_attribute_patterns = append(base64encode_attribute_patterns, base64encode_attribute_pattern.(string))
	}

	ldapEntry, err := cl.ReadEntryByFilter(ou, "("+filter+")", ignore_attributes, ignore_attribute_patterns, base64encode_attributes, base64encode_attribute_patterns)

	if err != nil {
		return diag.FromErr(err)
	}

	id := ldapEntry.Dn
	d.SetId(id)

	err = d.Set("dn", id)
	if err != nil {
		return diag.FromErr(err)
	}

	jsonData, err := json.Marshal(ldapEntry.Entry)
	if err != nil {
		return diag.Errorf("error marshaling JSON for %q: %s", id, err)
	}

	if err := d.Set("data_json", string(jsonData)); err != nil {
		return diag.FromErr(err)
	}

	return diag.FromErr(err)
}
