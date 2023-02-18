---
subcategory: "Elasticsearch"
layout: "aws"
page_title: "AWS: aws_elasticsearch_engine_version"
description: |-
  Get information on an Elasticsearch engine version.
---

# Data Source: aws_elasticsearch_domain

Use this data source to get information about Elasticsearch version

## Example Usage

```terraform
data "aws_elasticsearch_engine_version" "my_domain" {
  version = "OpenSearch_2.3"
}
```

## Argument Reference

The following arguments are supported:

* `version` – (Optional) Version of the Elasticsearch/Opensearch engine. Conflict with `preffered_versions`.
* `preferred_versions` - (Optional) Ordered list of preffered version. The first match in this list will be returned.
 Conflict with `version`

## Attributes Reference

The following attributes are exported:

* `version` – Elasticsearch/Opensearch engine version.
