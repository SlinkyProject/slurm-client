---
# yaml-language-server: $schema=https://raw.githubusercontent.com/oapi-codegen/oapi-codegen/{{OAPI_CODEGEN_VERSION}}/configuration-schema.json
package: {{SLURM_GO_MODULE}}
generate:
  models: true
output: object.gen.go
output-options:
  skip-fmt: false
  skip-prune: true
  user-templates:
    imports.tmpl: object-imports.tmpl
    typedef.tmpl: object.tmpl
    constants.tmpl: empty.tmpl
    param-types.tmpl: empty.tmpl
    request-bodies.tmpl: empty.tmpl
    additional-properties.tmpl: empty.tmpl
    union.tmpl: empty.tmpl
    union-and-additional-properties.tmpl: empty.tmpl
