---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_openapi_api Resource - terraform-provider-datadog"
subcategory: ""
description: |-
  Provides a Datadog OpenapiApi resource. This can be used to create and manage Datadog openapi_api.
---

# datadog_openapi_api (Resource)

Provides a Datadog OpenapiApi resource. This can be used to create and manage Datadog openapi_api.

## Example Usage

```terraform
# Create new openapi_api resource

resource "datadog_openapi_api" "my-api" {
  spec = file("./path/my-api.yaml")
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `spec` (String) The OpenAPI spec.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import datadog_openapi_api.new_list "90646597-5fdb-4a17-a240-647003f8c028"
```
