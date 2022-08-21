---
page_title: "ldap_entry Data Source - terraform-provider-ldap"
subcategory: ""
description: |-
---

# ldap_entry (Data Source)

Provides details about a user in LDAP. 

## Example Usage
```terraform
data "ldap_entry" "user" {
  ou     = "ou=People,dc=example,dc=com"
  filter = "mail=user@example.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `filter` (String) The filter for selecting the LDAP entry
- `ou` (String) OU where LDAP entry will be searched

### Read-Only

- `data_json` (String) JSON-encoded string that that is read as the attributes of the entry
- `id` (String) The DN of the LDAP entry