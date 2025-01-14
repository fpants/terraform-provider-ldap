---
page_title: "ldap_entry Data Source - terraform-provider-ldap"
subcategory: ""
description: |-
---

# ldap_entry (Data Source)

Provides details about an entry in LDAP. 

Attributes of the entry can be ignored by `ignore_attributes` or `ignore_attribute_patterns`.

Attributes of the entry can be encoded to base64 by `base64encode_attributes` or `base64encode_attribute_patterns`.
This should be used for attributes with binary content.

## Example Usage
```terraform
data "ldap_entry" "user" {
  ou     = "ou=People,dc=example,dc=com"
  filter = "mail=user@example.com"
}

data "ldap_entry" "user_jim_mit" {
  dn = "uid=jimmit01,ou=People,dc=example,dc=com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `base64encode_attribute_patterns` (List of String) list of attribute patterns to be encoded to base64
- `base64encode_attributes` (List of String) list of attributes to be encoded to base64
- `dn` (String) DN of the LDAP entry
- `filter` (String) filter for selecting the LDAP entry, ignored if 'dn' is used
- `ignore_attribute_patterns` (List of String) list of attribute patterns to ignore
- `ignore_attributes` (List of String) list of attributes to ignore
- `ou` (String) OU where LDAP entry will be searched

### Read-Only

- `data_json` (String) JSON-encoded string that is read as the values of the attributes of the entry (s. https://pkg.go.dev/github.com/go-ldap/ldap/v3#EntryAttribute)
- `id` (String) The ID of this resource.
